GOFILES = $(shell find . -name '*.go')

default: build

workdir:
	mkdir -p workdir

build: workdir/web_logger

build-native: $(GOFILES)
	go build -o workdir/native-web_logger .

workdir/guardian: $(GOFILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o workdir/web_logger .