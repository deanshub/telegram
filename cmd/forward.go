package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

func Forward(c *client.Client, args []string) {
	fs := flag.NewFlagSet("forward", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "destination chat ID (required)")
	fromChatID := fs.String("from-chat-id", "", "source chat ID (required)")
	messageID := fs.String("message-id", "", "message ID to forward (required)")
	silent := fs.Bool("silent", false, "send without notification")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" || *fromChatID == "" || *messageID == "" {
		fatal("forward requires --chat-id, --from-chat-id, and --message-id")
	}

	payload := map[string]any{
		"chat_id":      *chatID,
		"from_chat_id": *fromChatID,
		"message_id":   *messageID,
	}
	if *silent {
		payload["disable_notification"] = true
	}

	result, err := c.CallJSON("forwardMessage", payload)
	if err != nil {
		fatalf("forwardMessage: %v", err)
	}
	printJSON(result)
}
