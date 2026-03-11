---
name: tgbot
description: Send messages, media, and locations via Telegram using the tgbot CLI. Use when the user asks to send a Telegram message, check Telegram updates, or interact with a Telegram bot.
---

# tgbot

Minimal CLI for the Telegram Bot API. Requires `TELEGRAM_BOT_TOKEN` and optionally `TELEGRAM_CHAT_ID` environment variables.

## Commands

### Send a text message

```bash
tgbot send --chat-id <CHAT_ID> --text "Hello!"
```

Optional flags: `--parse-mode` (Markdown, MarkdownV2, HTML), `--silent`, `--reply-to <MSG_ID>`.

### Send media (photo, video, document, audio)

```bash
tgbot media --chat-id <CHAT_ID> --type photo --file /path/to/image.jpg
tgbot media --chat-id <CHAT_ID> --type video --file /path/to/video.mp4 --caption "Check this out"
```

To send a GIF, use `--type video` with a `.gif` file. Optional: `--caption`, `--parse-mode`, `--silent`.

### Get updates (poll for new messages)

```bash
tgbot updates
```

Returns a JSON array of update objects. Each contains `update_id`, `message.text`, `message.from`, and `message.chat.id`.

### Send a location

```bash
tgbot location --chat-id <CHAT_ID> --lat 32.0853 --lon 34.7818
```

### Forward a message

```bash
tgbot forward --chat-id <TO_CHAT_ID> --from-chat-id <FROM_CHAT_ID> --message-id <MSG_ID>
```

### Delete a message

```bash
tgbot delete --chat-id <CHAT_ID> --message-id <MSG_ID>
```

### Get bot info

```bash
tgbot get-me
```

### Get chat info

```bash
tgbot chat --chat-id <CHAT_ID>
```

## Notes

- All commands output JSON.
- Use `--token` flag or `TELEGRAM_BOT_TOKEN` env var for authentication.
- Use `--chat-id` flag or `TELEGRAM_CHAT_ID` env var for the default chat.
