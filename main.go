package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/dean/tgbot/client"
	"github.com/dean/tgbot/cmd"
)

const usage = `tgbot - Telegram Bot API CLI

Usage: tgbot [--token TOKEN] <command> [flags]

Commands:
  get-me     Get bot info
  send       Send a text message
  media      Send a photo/document/video/audio
  updates    Get updates (with optional long-polling)
  chat       Get chat info
  location   Send a location
  forward    Forward a message
  delete     Delete a message

Environment:
  TELEGRAM_BOT_TOKEN   bot token (or use --token flag)
  TELEGRAM_CHAT_ID     default chat ID (or use --chat-id flag)
`

func main() {
	args := os.Args[1:]
	token, args := extractToken(args)

	if token == "" {
		token = os.Getenv("TELEGRAM_BOT_TOKEN")
	}

	if len(args) == 0 {
		fmt.Print(usage)
		os.Exit(0)
	}

	command := args[0]
	cmdArgs := args[1:]

	if command == "help" || command == "--help" || command == "-h" {
		fmt.Print(usage)
		os.Exit(0)
	}

	if token == "" {
		fmt.Fprintln(os.Stderr, "error: token required (use --token or TELEGRAM_BOT_TOKEN env var)")
		os.Exit(1)
	}

	c := client.New(token)

	switch command {
	case "get-me":
		cmd.GetMe(c, cmdArgs)
	case "send":
		cmd.Send(c, cmdArgs)
	case "media":
		cmd.Media(c, cmdArgs)
	case "updates":
		cmd.Updates(c, cmdArgs)
	case "chat":
		cmd.ChatInfo(c, cmdArgs)
	case "location":
		cmd.Location(c, cmdArgs)
	case "forward":
		cmd.Forward(c, cmdArgs)
	case "delete":
		cmd.Delete(c, cmdArgs)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", command)
		fmt.Print(usage)
		os.Exit(1)
	}
}

// extractToken scans args for --token VALUE or --token=VALUE, returns the token
// and the remaining args with the token flag removed.
func extractToken(args []string) (string, []string) {
	var token string
	var remaining []string

	for i := 0; i < len(args); i++ {
		arg := args[i]
		if arg == "--token" && i+1 < len(args) {
			token = args[i+1]
			i++ // skip next
		} else if strings.HasPrefix(arg, "--token=") {
			token = strings.TrimPrefix(arg, "--token=")
		} else {
			remaining = append(remaining, arg)
		}
	}
	return token, remaining
}
