package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

var mediaTypes = map[string]struct {
	method string
	field  string
}{
	"photo":    {"sendPhoto", "photo"},
	"document": {"sendDocument", "document"},
	"video":    {"sendVideo", "video"},
	"audio":    {"sendAudio", "audio"},
}

func Media(c *client.Client, args []string) {
	fs := flag.NewFlagSet("media", flag.ExitOnError)
	chatID := fs.String("chat-id", "", "target chat ID (required)")
	mediaType := fs.String("type", "", "media type: photo, document, video, audio (required)")
	file := fs.String("file", "", "path to file (required)")
	caption := fs.String("caption", "", "media caption")
	parseMode := fs.String("parse-mode", "", "caption parse mode: Markdown, MarkdownV2, HTML")
	silent := fs.Bool("silent", false, "send without notification")
	fs.Parse(args)
	defaultChatID(chatID)

	if *chatID == "" || *mediaType == "" || *file == "" {
		fatal("media requires --chat-id, --type, and --file")
	}

	mt, ok := mediaTypes[*mediaType]
	if !ok {
		fatal("unknown media type: " + *mediaType + " (use photo, document, video, audio)")
	}

	fields := map[string]string{
		"chat_id": *chatID,
	}
	if *caption != "" {
		fields["caption"] = *caption
	}
	if *parseMode != "" {
		fields["parse_mode"] = *parseMode
	}
	if *silent {
		fields["disable_notification"] = "true"
	}

	result, err := c.UploadFile(mt.method, mt.field, *file, fields)
	if err != nil {
		fatalf("%s: %v", mt.method, err)
	}
	printJSON(result)
}
