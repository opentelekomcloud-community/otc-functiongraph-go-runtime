mkfile_path := $(abspath $(lastword $(MAKEFILE_LIST)))/..
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))

BACKEND_CONFIG_BUCKET := "doc-samples-tf-backend"
BACKEND_CONFIG_KEY := "terraform_state/go/doc-sample-event-s3obs-thumbnail.tf"
BACKEND_CONFIG_REGION := "eu-de"
BACKEND_CONFIG_ENDPOINTS := "endpoints={s3=\"https://obs.eu-de.otc.t-systems.com\"}"
HANDLER_NAME=example

clean:
	rm -f $(mkfile_path)/deploy.zip

build:
	GOARCH=amd64 GO111MODULE=on GOOS=linux CGO_ENABLED=0 go build -o target/${HANDLER_NAME} src/main.go

all: build

release: all
	printf "# Handler name\n${HANDLER_NAME}" > target/handler_name.txt
	cd target && zip -r ../deploy.zip .

tf_init:
	terraform -chdir=terraform \
	  init \
	  -backend-config=$(BACKEND_CONFIG_ENDPOINTS) \
	  -backend-config="bucket=$(BACKEND_CONFIG_BUCKET)" \
	  -backend-config="key=$(BACKEND_CONFIG_KEY)" \
	  -backend-config="region=$(BACKEND_CONFIG_REGION)"

tf_plan: release
	if [ ! -f "terraform/.terraform.lock.hcl" ]; then \
		$(MAKE) initTerraform; \
	fi
	terraform -chdir=terraform \
	  plan \
		-var-file="variables.tfvars"

tf_apply: release
	if [ ! -f "terraform/.terraform.lock.hcl" ]; then \
		$(MAKE) initTerraform; \
	fi
	terraform -chdir=terraform \
	  apply -auto-approve \
	  -var-file="variables.tfvars"

tf_destroy:
	terraform -chdir=terraform \
	  destroy -auto-approve \
		-var-file="variables.tfvars"

.PHONY: build zip clean all release tf_init tf_plan tf_apply tf_destroy
