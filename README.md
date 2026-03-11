# tgbot

Minimal CLI for the Telegram Bot API. Zero dependencies, single binary.

## Install

Download a binary from [Releases](../../releases), or build from source:

```sh
go install github.com/dean/tgbot@latest
```

Or clone and build:

```sh
git clone https://github.com/dean/tgbot.git
cd tgbot
make build-small
```

## Configuration

| Environment Variable | Description |
|---|---|
| `TELEGRAM_BOT_TOKEN` | Bot token (or use `--token` flag) |
| `TELEGRAM_CHAT_ID` | Default chat ID (or use `--chat-id` flag) |

## Usage

```sh
# Bot info
tgbot get-me

# Send a message
tgbot send --text "hello world"

# Send with Markdown
tgbot send --chat-id 123456 --text "*bold*" --parse-mode Markdown

# Send a photo
tgbot media --type photo --file photo.jpg --caption "nice pic"

# Send a document/video/audio
tgbot media --type video --file clip.mp4

# Send a location
tgbot location --lat 51.5074 --lon -0.1278

# Forward a message
tgbot forward --from-chat-id 111 --message-id 42

# Delete a message
tgbot delete --message-id 42

# Get chat info
tgbot chat --chat-id 123456

# Fetch recent updates
tgbot updates

# Long-poll for updates (Ctrl+C to stop)
tgbot updates --poll
```

All commands output JSON, pipeable to `jq`:

```sh
tgbot get-me | jq .username
```

## Commands

| Command | Description |
|---|---|
| `get-me` | Get bot info |
| `send` | Send a text message |
| `media` | Send a photo, document, video, or audio |
| `updates` | Get updates or long-poll with `--poll` |
| `chat` | Get chat info |
| `location` | Send a location |
| `forward` | Forward a message |
| `delete` | Delete a message |

## Build targets

```sh
make build         # standard build
make build-small   # stripped, static binary (~5MB)
make build-tiny    # + UPX compressed (~1.5MB, requires upx)
make build-linux   # cross-compile linux/amd64
make build-linux-arm  # cross-compile linux/arm64
```
