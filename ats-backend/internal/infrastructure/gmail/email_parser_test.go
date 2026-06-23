package gmail

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/api/gmail/v1"
)

func TestParseEmail_Basic(t *testing.T) {
	msg := &gmail.Message{
		Id:       "msg-123",
		ThreadId: "thread-123",
		Payload: &gmail.MessagePart{
			Headers: []*gmail.MessagePartHeader{
				{Name: "From", Value: "John Doe <john@example.com>"},
				{Name: "Subject", Value: "Resume Attached"},
			},
			Body: &gmail.MessagePartBody{
				Data: base64.URLEncoding.EncodeToString([]byte("Hello, here is my resume.")),
			},
		},
	}

	// We pass nil for the service since there are no attachments requiring API fetches
	parsed, err := ParseEmail(msg, nil)

	assert.NoError(t, err)
	assert.Equal(t, "msg-123", parsed.MessageID)
	assert.Equal(t, "John Doe", parsed.FromName)
	assert.Equal(t, "john@example.com", parsed.FromEmail)
	assert.Equal(t, "Resume Attached", parsed.Subject)
	assert.Equal(t, "Hello, here is my resume.", parsed.BodyText)
}
