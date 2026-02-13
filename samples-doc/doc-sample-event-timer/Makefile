HANDLER_NAME=timer-example

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) src/main.go

zip:
	printf "# Handler name\n${HANDLER_NAME}" > target/handler_name.txt
	cd target && zip -r ../deploy.zip . 

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip

tf_init:
	terraform -chdir=terraform \
	  init \
	  -backend-config="endpoints={s3=\"https://obs.eu-de.otc.t-systems.com\"}" \
	  -backend-config="bucket=sample-tf-backend" \
	  -backend-config="key=terraform_state/go/doc-sample-event-timer.tf" \
	  -backend-config="region=eu-de"

tf_plan: 
	if [ ! -f "terraform/.terraform.lock.hcl" ]; then \
		$(MAKE) tf_init; \
	fi
	terraform -chdir=terraform \
	  plan \
	  -var function_name="doc-sample-event-timer" \
	  -var function_handler_name=$(HANDLER_NAME)

tf_apply: all
	if [ ! -f "terraform/.terraform.lock.hcl" ]; then \
		$(MAKE) tf_init; \
	fi
	terraform -chdir=terraform \
	  apply -auto-approve \
	  -var function_name="doc-sample-event-timer" \
	  -var function_handler_name=$(HANDLER_NAME)

tf_destroy:
	terraform -chdir=terraform \
	  destroy -auto-approve

.PHONY: build zip clean all tf_init tf_plan tf_apply tf_destroy
