package chat

import (
	"fmt"

	"github.com/twilio/twilio-go"
	conversations "github.com/twilio/twilio-go/rest/conversations/v1"
)

var client = twilio.NewRestClient()

func CreateConversation(friendlyName string) (string, error) {
    params := &conversations.CreateConversationParams{}
    params.SetFriendlyName(friendlyName)

    resp, err := client.ConversationsV1.CreateConversation(params)
    if err != nil {
        return "", fmt.Errorf("failed to create conversation: %w", err)
    }

    return *resp.Sid, nil
}

func AddParticipant(conversationSid, identity string) error {
    params := &conversations.CreateConversationParticipantParams{}
    params.SetIdentity(identity)

    _, err := client.ConversationsV1.CreateConversationParticipant(conversationSid, params)
    if err != nil {
        return fmt.Errorf("failed to add participant: %w", err)
    }

    return nil
}

func SendMessage(conversationSid, author, body string) error {
    params := &conversations.CreateConversationMessageParams{}
    params.SetAuthor(author)
    params.SetBody(body)

    _, err := client.ConversationsV1.CreateConversationMessage(conversationSid, params)
    if err != nil {
        return fmt.Errorf("failed to send message: %w", err)
    }

    return nil
}