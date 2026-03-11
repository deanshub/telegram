package cmd

import (
	"flag"

	"github.com/dean/tgbot/client"
)

func GetMe(c *client.Client, args []string) {
	fs := flag.NewFlagSet("get-me", flag.ExitOnError)
	fs.Parse(args)

	result, err := c.Call("getMe", nil)
	if err != nil {
		fatalf("getMe: %v", err)
	}
	printJSON(result)
}
