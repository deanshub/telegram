package cmd

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"strconv"

	"github.com/dean/tgbot/client"
)

func Updates(c *client.Client, args []string) {
	fs := flag.NewFlagSet("updates", flag.ExitOnError)
	poll := fs.Bool("poll", false, "long-poll for updates continuously")
	timeout := fs.Int("timeout", 30, "long-poll timeout in seconds")
	limit := fs.Int("limit", 100, "max updates per request")
	offset := fs.Int64("offset", 0, "update offset")
	fs.Parse(args)

	if !*poll {
		params := url.Values{}
		if *limit > 0 {
			params.Set("limit", strconv.Itoa(*limit))
		}
		if *offset > 0 {
			params.Set("offset", strconv.FormatInt(*offset, 10))
		}
		result, err := c.Call("getUpdates", params)
		if err != nil {
			fatalf("getUpdates: %v", err)
		}
		printJSON(result)
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	currentOffset := *offset
	for {
		select {
		case <-ctx.Done():
			fmt.Fprintln(os.Stderr, "\nstopping poll...")
			return
		default:
		}

		params := url.Values{
			"timeout": {strconv.Itoa(*timeout)},
			"limit":   {strconv.Itoa(*limit)},
		}
		if currentOffset > 0 {
			params.Set("offset", strconv.FormatInt(currentOffset, 10))
		}

		result, err := c.Call("getUpdates", params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "poll error: %v\n", err)
			continue
		}

		var updates []client.Update
		if err := json.Unmarshal(result, &updates); err != nil {
			fmt.Fprintf(os.Stderr, "decode error: %v\n", err)
			continue
		}

		for _, u := range updates {
			raw, _ := json.MarshalIndent(u, "", "  ")
			fmt.Println(string(raw))
			if u.UpdateID >= currentOffset {
				currentOffset = u.UpdateID + 1
			}
		}
	}
}
