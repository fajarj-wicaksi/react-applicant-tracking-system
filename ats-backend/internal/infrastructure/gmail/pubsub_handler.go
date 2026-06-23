package gmail

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PubSubMessage struct {
	Message struct {
		Data        []byte `json:"data"` // Base64 encoded JSON
		MessageId   string `json:"messageId"`
		PublishTime string `json:"publishTime"`
	} `json:"message"`
}

type GmailPushData struct {
	EmailAddress string `json:"emailAddress"`
	HistoryId    uint64 `json:"historyId"`
}

// HandleGmailWebhook handles the incoming push from Google Cloud Pub/Sub
func HandleGmailWebhook(c *gin.Context, enqueueFunc func(emailAddress string, historyId uint64)) {
	var payload PubSubMessage
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pubsub payload"})
		return
	}

	var data GmailPushData
	if err := json.Unmarshal(payload.Message.Data, &data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid base64 data"})
		return
	}

	// Always return 200 OK to acknowledge Pub/Sub
	c.Status(http.StatusOK)

	// Async queue processing
	go enqueueFunc(data.EmailAddress, data.HistoryId)
}
