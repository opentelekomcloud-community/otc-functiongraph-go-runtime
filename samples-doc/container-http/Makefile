TARGET_PATH=target
#DOCKER_FILE=Dockerfile
DOCKER_FILE=Dockerfile.alpine

IMAGE_NAME=custom_container_http_example

build: clean
	mkdir -p $(TARGET_PATH)
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 && \
	cd src && \
	go build -o ../$(TARGET_PATH)/httpfunction .

run_local:
	cd $(TARGET_PATH) && ./httpfunction

docker_build: build
	docker buildx build \
	  --platform linux/amd64 \
		--build-arg FILE_PATH=$(TARGET_PATH) \
		--file $(DOCKER_FILE) \
		--tag $(IMAGE_NAME):latest .

docker_run_local: docker_build
	docker container run \
	  --rm \
		--platform linux/amd64 \
		--publish 8000:8000 \
		--name container_http_example \
		$(IMAGE_NAME):latest

docker_push: docker_build
  # see: https://docs.otc.t-systems.com/software-repository-container/umn/image_management/obtaining_a_long-term_valid_login_command.html#swr-01-1000
	docker login -u $(OTC_SDK_PROJECTNAME)@$(OTC_SDK_AK) -p $(OTC_SWR_LOGIN_KEY) $(OTC_SWR_ENDPOINT)
	docker tag $(IMAGE_NAME):latest $(OTC_SWR_ENDPOINT)/$(OTC_SWR_ORGANIZATION)/$(IMAGE_NAME):latest
	docker push $(OTC_SWR_ENDPOINT)/$(OTC_SWR_ORGANIZATION)/$(IMAGE_NAME):latest

docker_all: build docker_build docker_push

test_local:
  # execute a curl 
	@curl -X POST \
	  -H "X-Cff-Request-Id: $(shell uuidgen)" \
	  -H "X-Cff-Func-Name: 0@default@sample-container-http" \
		-H "X-Cff-Func-Version: latest" \
	  -H 'Content-Type: application/json' \		
		-d '{"message": "Hello from API Gateway POST request"}' \
		http://localhost:8000/index
	@echo ""

clean:
	rm -rf $(TARGET_PATH)

all: build docker_build

.PHONY: build run_local docker_build docker_run_local docker_push docker_all test_local clean
