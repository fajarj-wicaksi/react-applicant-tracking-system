package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/cenkalti/backoff/v4"
	"google.golang.org/api/gmail/v1"

	infragmail "ats-backend/internal/infrastructure/gmail"
	inframinio "ats-backend/internal/infrastructure/minio"
)

type EmailIngestService struct {
	oauthClient   *infragmail.OAuthClient
	storageClient *inframinio.StorageClient
	// In reality, we'd inject repositories here for DB access
	// candidateRepo port.CandidateRepository
	// emailRepo     port.EmailRepository
}

func NewEmailIngestService(oauth *infragmail.OAuthClient, storage *inframinio.StorageClient) *EmailIngestService {
	return &EmailIngestService{
		oauthClient:   oauth,
		storageClient: storage,
	}
}

// ProcessHistory changes for a given user account
func (s *EmailIngestService) ProcessHistory(ctx context.Context, accountID string, startHistoryId uint64) error {
	srv, err := s.oauthClient.GetGmailService(ctx, accountID)
	if err != nil {
		return err
	}

	// Use backoff to handle rate limits (429)
	var historyResp *gmail.ListHistoryResponse
	operation := func() error {
		req := srv.Users.History.List("me").StartHistoryId(startHistoryId)
		resp, err := req.Do()
		if err != nil {
			return err
		}
		historyResp = resp
		return nil
	}

	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = 2 * time.Minute

	if err := backoff.Retry(operation, b); err != nil {
		return fmt.Errorf("history fetch failed after retries: %w", err)
	}

	for _, history := range historyResp.History {
		for _, msgAdded := range history.MessagesAdded {
			err := s.processMessage(ctx, srv, msgAdded.Message.Id)
			if err != nil {
				log.Printf("failed to process message %s: %v", msgAdded.Message.Id, err)
			}
		}
	}

	return nil
}

func (s *EmailIngestService) processMessage(ctx context.Context, srv *gmail.Service, messageID string) error {
	msg, err := srv.Users.Messages.Get("me", messageID).Format("full").Do()
	if err != nil {
		return err
	}

	parsed, err := infragmail.ParseEmail(msg, srv)
	if err != nil {
		return err
	}

	// 1. Upload Attachments to MinIO
	var resumeKey string
	for _, attachment := range parsed.Attachments {
		objectKey := fmt.Sprintf("candidates/%s/%s", messageID, attachment.Filename)
		_, err := s.storageClient.UploadFile(ctx, objectKey, attachment.Data, attachment.MimeType)
		if err != nil {
			log.Printf("Warning: failed to upload attachment: %v", err)
			continue
		}
		if resumeKey == "" && attachment.MimeType == "application/pdf" {
			resumeKey = objectKey // naive logic: first PDF is resume
		}
	}

	// 2. Candidate Creation Logic (Pseudo-code as we lack repos)
	// candidate, err := s.candidateRepo.FindByEmail(ctx, parsed.FromEmail)
	// if candidate == nil {
	//    candidate = &domain.Candidate{
	//        FirstName: parsed.FromName,
	//        Email: parsed.FromEmail,
	//        ResumeUrl: resumeKey,
	//    }
	//    s.candidateRepo.Create(ctx, candidate)
	// }
	
	// 3. Email Activity Record
	// s.emailRepo.Create(ctx, &domain.Email{
	//    MessageId: parsed.MessageID,
	//    Subject: parsed.Subject,
	//    Body: parsed.BodyText,
	//    CandidateId: candidate.ID,
	// })

	log.Printf("Successfully ingested email from %s with %d attachments", parsed.FromEmail, len(parsed.Attachments))
	return nil
}
