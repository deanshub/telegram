package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

func Location(c *client.Client, args []string) {
	fs := flag.NewFlagSet("location", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "target chat ID (required)")
	lat := fs.Float64("lat", 0, "latitude (required)")
	lon := fs.Float64("lon", 0, "longitude (required)")
	silent := fs.Bool("silent", false, "send without notification")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" {
		fatal("location requires --chat-id")
	}

	payload := map[string]any{
		"chat_id":   *chatID,
		"latitude":  *lat,
		"longitude": *lon,
	}
	if *silent {
		payload["disable_notification"] = true
	}

	result, err := c.CallJSON("sendLocation", payload)
	if err != nil {
		fatalf("sendLocation: %v", err)
	}
	printJSON(result)
}
