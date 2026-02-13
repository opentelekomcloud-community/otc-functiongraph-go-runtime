HANDLER_NAME=go-http-demo

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) src/main.go

zip:
	# Create bootstrap file pointing to the handler binary  
	@echo "/opt/function/code/$(HANDLER_NAME)" > target/bootstrap
	
	# Zip all files in target directory
	cd target && zip -r ../deploy.zip .

clean:
	rm -f deploy.zip
	rm -rf target

test_local_post:
	@curl -X POST \
	-H 'Content-Type: application/json' \
	-d 'nice to meet you local' \
	http://localhost:8000/hello

test_local_get:
	@curl -X GET \
	http://localhost:8000/hello 

all: clean build zip

.PHONY: build zip clean all test_local_post test_local_get