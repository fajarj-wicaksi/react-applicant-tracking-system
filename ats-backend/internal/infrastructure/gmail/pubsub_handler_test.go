package gmail

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleGmailWebhook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	var enqueuedEmail string
	var enqueuedHistoryId uint64

	mockEnqueue := func(emailAddress string, historyId uint64) {
		enqueuedEmail = emailAddress
		enqueuedHistoryId = historyId
	}

	router.POST("/webhook", func(c *gin.Context) {
		HandleGmailWebhook(c, mockEnqueue)
	})

	pushData := GmailPushData{
		EmailAddress: "test@example.com",
		HistoryId:    12345,
	}
	pushBytes, _ := json.Marshal(pushData)

	payload := PubSubMessage{}
	payload.Message.Data = pushBytes

	payloadBytes, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/webhook", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	// Wait a tiny bit for the goroutine
	time.Sleep(100 * time.Millisecond)
	assert.Equal(t, "test@example.com", enqueuedEmail)
	assert.Equal(t, uint64(12345), enqueuedHistoryId)
}
