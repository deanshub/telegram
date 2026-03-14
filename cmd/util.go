package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printJSON(data json.RawMessage) {
	var buf []byte
	buf, err := json.MarshalIndent(json.RawMessage(data), "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error formatting json: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(buf))
}

func defaultChatID(flagVal *string) {
	if *flagVal == "" {
		*flagVal = os.Getenv("TELEGRAM_CHAT_ID")
	}
}

func parseAllowedIDs(flagVal string) map[int64]bool {
	raw := flagVal
	if raw == "" {
		raw = os.Getenv("TELEGRAM_ALLOWED_USER_IDS")
	}
	if raw == "" {
		raw = os.Getenv("TELEGRAM_CHAT_ID")
	}
	if raw == "" {
		return nil
	}
	allowed := make(map[int64]bool)
	for _, s := range strings.Split(raw, ",") {
		s = strings.TrimSpace(s)
		if id, err := strconv.ParseInt(s, 10, 64); err == nil {
			allowed[id] = true
		}
	}
	if len(allowed) == 0 {
		return nil
	}
	return allowed
}

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
