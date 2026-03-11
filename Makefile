BINARY = tgbot
LDFLAGS = -ldflags="-s -w"

.PHONY: build build-small build-tiny clean

build:
	go build -o $(BINARY) .

build-small:
	CGO_ENABLED=0 go build -trimpath $(LDFLAGS) -o $(BINARY) .

build-tiny: build-small
	upx --best --lzma $(BINARY)

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath $(LDFLAGS) -o $(BINARY)-linux-amd64 .

build-linux-arm:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -trimpath $(LDFLAGS) -o $(BINARY)-linux-arm64 .

clean:
	rm -f $(BINARY) $(BINARY)-linux-*
