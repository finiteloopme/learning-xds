APP_NAME=xds-from-scratch
MOD_FQ_NAME=github.com/finiteloopme/xds-from-scratch
OUTPUT_FOLDER=bin
BINARY_NAME=xds-from-scratch
CMD_NAME=cmd
GCP_REGION=us-central1
GCP_ZONE=${GCP_REGION}-a

create-module:
	go mod init ${MOD_FQ_NAME}

fmt-deps:
	# go vet
	# go fmt
	go mod tidy

build: fmt-deps
	go build -o ${OUTPUT_FOLDER}/server ./${CMD_NAME}/server/...
	go build -o ${OUTPUT_FOLDER}/client ./${CMD_NAME}/client/...
	# go build ./pkg/...

run: 
	go run ${CMD_NAME}/main.go

grpc-mod-update:
	cd api; buf mod update

grpc-mod-lint: grpc-mod-update
	cd api; buf lint

grpc-codegen: grpc-mod-update
	buf generate

# Ensure `create-module` target has been executed once
cloud-run:
	gcloud run deploy ${APP_NAME} --platform=managed --allow-unauthenticated --region=${GCP_REGION} --source=.

undeploy-cloud-run:
	gcloud run services delete new-cr-go-module --platform=managed --region=us-central1 --quiet