package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

func Send(c *client.Client, args []string) {
	fs := flag.NewFlagSet("send", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "target chat ID (required)")
	text := fs.String("text", "", "message text (required)")
	parseMode := fs.String("parse-mode", "", "parse mode: Markdown, MarkdownV2, HTML")
	silent := fs.Bool("silent", false, "send without notification")
	replyTo := fs.String("reply-to", "", "message ID to reply to")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" || *text == "" {
		fatal("send requires --chat-id and --text")
	}

	payload := map[string]any{
		"chat_id": *chatID,
		"text":    *text,
	}
	if *parseMode != "" {
		payload["parse_mode"] = *parseMode
	}
	if *silent {
		payload["disable_notification"] = true
	}
	if *replyTo != "" {
		payload["reply_to_message_id"] = *replyTo
	}

	result, err := c.CallJSON("sendMessage", payload)
	if err != nil {
		fatalf("sendMessage: %v", err)
	}
	printJSON(result)
}
