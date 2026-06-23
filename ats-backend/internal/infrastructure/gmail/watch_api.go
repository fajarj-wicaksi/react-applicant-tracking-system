package gmail

import (
	"fmt"

	"google.golang.org/api/gmail/v1"
)

type WatchManager struct{}

func NewWatchManager() *WatchManager {
	return &WatchManager{}
}

// StartWatch configures the Gmail account to send push notifications to a Pub/Sub topic
func (w *WatchManager) StartWatch(srv *gmail.Service, topicName string) (*gmail.WatchResponse, error) {
	req := &gmail.WatchRequest{
		LabelIds:  []string{"INBOX"},
		TopicName: topicName,
	}

	resp, err := srv.Users.Watch("me", req).Do()
	if err != nil {
		return nil, fmt.Errorf("failed to start watch: %w", err)
	}

	return resp, nil
}

// StopWatch stops push notifications
func (w *WatchManager) StopWatch(srv *gmail.Service) error {
	err := srv.Users.Stop("me").Do()
	if err != nil {
		return fmt.Errorf("failed to stop watch: %w", err)
	}
	return nil
}
