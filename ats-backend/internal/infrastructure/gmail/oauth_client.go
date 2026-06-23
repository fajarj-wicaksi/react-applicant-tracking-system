package gmail

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// TokenStore defines how we retrieve and save tokens from the database.
type TokenStore interface {
	GetToken(ctx context.Context, accountID string) (*oauth2.Token, error)
	SaveToken(ctx context.Context, accountID string, token *oauth2.Token) error
}

type OAuthClient struct {
	config     *oauth2.Config
	tokenStore TokenStore
}

func NewOAuthClient(clientID, clientSecret, redirectURL string, store TokenStore) *OAuthClient {
	return &OAuthClient{
		config: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{gmail.GmailReadonlyScope},
			Endpoint:     google.Endpoint,
		},
		tokenStore: store,
	}
}

// GetGmailService returns a ready-to-use Gmail API service client.
// It automatically handles token refreshes using the provided token source.
func (o *OAuthClient) GetGmailService(ctx context.Context, accountID string) (*gmail.Service, error) {
	token, err := o.tokenStore.GetToken(ctx, accountID)
	if err != nil {
		return nil, err
	}

	tokenSource := o.config.TokenSource(ctx, token)
	client := oauth2.NewClient(ctx, tokenSource)
	return gmail.NewService(ctx, option.WithHTTPClient(client))
}
