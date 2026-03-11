package cmd

import (
	"encoding/json"
	"fmt"
	"os"
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

func fatal(msg string) {
	fmt.Fprintln(os.Stderr, "error:", msg)
	os.Exit(1)
}

func fatalf(format string, args ...any) {
	fmt.Fprintf(os.Stderr, "error: "+format+"\n", args...)
	os.Exit(1)
}
