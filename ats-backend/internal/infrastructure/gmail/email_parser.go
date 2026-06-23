package gmail

import (
	"encoding/base64"
	"fmt"
	"strings"

	"google.golang.org/api/gmail/v1"
)

type ParsedEmail struct {
	MessageID   string
	ThreadID    string
	FromEmail   string
	FromName    string
	Subject     string
	BodyText    string
	Attachments []Attachment
}

type Attachment struct {
	Filename    string
	MimeType    string
	Data        []byte
	AttachmentID string
}

func ParseEmail(msg *gmail.Message, srv *gmail.Service) (*ParsedEmail, error) {
	parsed := &ParsedEmail{
		MessageID: msg.Id,
		ThreadID:  msg.ThreadId,
	}

	// Parse Headers
	for _, header := range msg.Payload.Headers {
		if header.Name == "From" {
			parsed.FromName, parsed.FromEmail = extractNameAndEmail(header.Value)
		}
		if header.Name == "Subject" {
			parsed.Subject = header.Value
		}
	}

	// Recursively parse parts
	if msg.Payload.Parts != nil {
		err := parseParts(msg.Payload.Parts, parsed, msg.Id, srv)
		if err != nil {
			return nil, err
		}
	} else if msg.Payload.Body != nil && msg.Payload.Body.Data != "" {
		// No parts, simple email
		decoded, _ := base64.URLEncoding.DecodeString(msg.Payload.Body.Data)
		parsed.BodyText = string(decoded)
	}

	return parsed, nil
}

func parseParts(parts []*gmail.MessagePart, parsed *ParsedEmail, msgId string, srv *gmail.Service) error {
	for _, part := range parts {
		if part.Filename != "" && part.Body.AttachmentId != "" {
			// It's an attachment
			attachReq := srv.Users.Messages.Attachments.Get("me", msgId, part.Body.AttachmentId)
			attach, err := attachReq.Do()
			if err != nil {
				return fmt.Errorf("failed to get attachment %s: %w", part.Filename, err)
			}
			
			decoded, err := base64.URLEncoding.DecodeString(attach.Data)
			if err != nil {
				continue
			}

			parsed.Attachments = append(parsed.Attachments, Attachment{
				Filename:     part.Filename,
				MimeType:     part.MimeType,
				Data:         decoded,
				AttachmentID: part.Body.AttachmentId,
			})
		} else if part.MimeType == "text/plain" && parsed.BodyText == "" {
			decoded, _ := base64.URLEncoding.DecodeString(part.Body.Data)
			parsed.BodyText = string(decoded)
		} else if part.Parts != nil {
			// Recursive call for nested multipart
			parseParts(part.Parts, parsed, msgId, srv)
		}
	}
	return nil
}

func extractNameAndEmail(from string) (string, string) {
	// Simple extraction: "John Doe <john@example.com>" -> "John Doe", "john@example.com"
	if idx := strings.Index(from, "<"); idx != -1 {
		name := strings.TrimSpace(from[:idx])
		name = strings.Trim(name, `"`)
		email := strings.Trim(from[idx+1:], ">")
		return name, email
	}
	return "", from // Fallback
}
