HANDLER_NAME=helloworld

build:
	go build -o target/$(HANDLER_NAME) main.go

zip:
	zip -j deploy.zip target/$(HANDLER_NAME)

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip
