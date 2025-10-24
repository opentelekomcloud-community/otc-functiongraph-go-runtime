HANDLER_NAME=timer-example

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -o target/$(HANDLER_NAME) src/main.go

zip:
	zip -j deploy.zip target/$(HANDLER_NAME)

clean:
	rm -f deploy.zip
	rm -rf target

all: build zip

initTerraform:
	terraform -chdir=terraform \
	  init \
	  -backend-config="endpoints={s3=\"https://obs.eu-de.otc.t-systems.com\"}" \
	  -backend-config="bucket=sample-tf-backend" \
	  -backend-config="key=terraform_state/go/doc-sample-event-timer.tf" \
	  -backend-config="region=eu-de"

deploy: all
	if [ ! -f "terraform/.terraform.lock.hcl" ]; then \
		$(MAKE) initTerraform; \
	fi
	terraform -chdir=terraform \
	  apply -auto-approve \
	  -var function_name="doc-sample-event-timer" \
	  -var function_handler_name=$(HANDLER_NAME)

destroy:
	terraform -chdir=terraform \
	  destroy -auto-approve

.PHONY: build zip clean all deploy destroy initTerraform
