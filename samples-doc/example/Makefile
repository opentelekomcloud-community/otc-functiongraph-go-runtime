

HANDLER_NAME=example

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) main.go

zip:
	zip -j deploy.zip target/$(HANDLER_NAME)

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip