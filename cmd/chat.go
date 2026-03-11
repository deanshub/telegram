package cmd

import (
	"flag"
	"net/url"

	"github.com/dean/tgbot/client"
)

func ChatInfo(c *client.Client, args []string) {
	fs := flag.NewFlagSet("chat", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "chat ID (required)")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" {
		fatal("chat requires --chat-id")
	}

	result, err := c.Call("getChat", url.Values{"chat_id": {*chatID}})
	if err != nil {
		fatalf("getChat: %v", err)
	}
	printJSON(result)
}
