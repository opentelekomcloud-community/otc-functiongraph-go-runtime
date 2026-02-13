HANDLER_NAME=example

build: clean
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) src/main.go

zip: build
	printf "# Handler name\n${HANDLER_NAME}" > target/handler_name.txt
	cd target && zip -r ../deploy.zip .

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip

.PHONY: build zip clean all