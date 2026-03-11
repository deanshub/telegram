package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

func Delete(c *client.Client, args []string) {
	fs := flag.NewFlagSet("delete", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "chat ID (required)")
	messageID := fs.String("message-id", "", "message ID to delete (required)")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" || *messageID == "" {
		fatal("delete requires --chat-id and --message-id")
	}

	payload := map[string]any{
		"chat_id":    *chatID,
		"message_id": *messageID,
	}

	result, err := c.CallJSON("deleteMessage", payload)
	if err != nil {
		fatalf("deleteMessage: %v", err)
	}
	printJSON(result)
}
