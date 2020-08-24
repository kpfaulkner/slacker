package slacker

import (
	"github.com/slack-go/slack"
)

type SlackHandler struct {
	slackKey string

	// general slack client
	client *slack.Client

	// realtime messaging
	rtm *slack.RTM
}

// NewSlackHandler creates new SlackHandler that will be a client to interact against the Slack APIs
// This should be a fairly light weight wrapper against the slack-go/slack client.
func NewSlackHandler(key string) *SlackHandler {
	sh := SlackHandler{}
	sh.slackKey = key
	sh.client = slack.New(sh.slackKey)
	sh.rtm = sh.client.NewRTM()

	return &sh
}

func (sh *SlackHandler) GetContacts() ([]Contact, error) {
	return nil, nil
}

func (sh *SlackHandler) PostMessage(message string, destination string) error {

	return nil
}

// IncomingMessageLoop loop permanently, receiving messages and posting them off places.
func (sh *SlackHandler) IncomingMessageLoop() error {

	return nil
}
