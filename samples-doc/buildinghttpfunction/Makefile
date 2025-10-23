HANDLER_NAME=go-http-demo

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) main.go

zip:
	# Create bootstrap file pointing to the handler binary  
	@echo "/opt/function/code/$(HANDLER_NAME)" > target/bootstrap
	
	# Zip all files in target directory
	zip -j deploy.zip target/**

clean:
	rm -f deploy.zip
	rm -rf target

all: clean build zip