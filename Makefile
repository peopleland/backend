USER_API_PATH=api/user/v1
USER_API_PROTO_FILES=$(shell cd $(USER_API_PATH) && find . -name *.proto)

install:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/bqx619/protoc-gen-go-lambda@latest
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

init:
	go mod download

generate:
	go generate ./...

lambda:
	cd $(USER_API_PATH) && protoc --proto_path=. \
		--proto_path=../../../third_party \
		--go_out=paths=source_relative:. \
        --go-lambda_out=paths=source_relative,prefix=.netlify/functions:. \
        --openapiv2_out ../../../website/swagger \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt json_names_for_fields=false \
        --openapiv2_opt disable_default_errors=true \
		$(USER_API_PROTO_FILES)

test:
	go test -v ./... -cover

build-user:
	mkdir -p bin/ && go build -o bin/user app/user/cmd/main.go && go build -o bin/sync-opener-record-background app/user/cmd/background/syncopenerrecordbackground.go && go build -o website/syncmonitor app/user/cmd/syncmonitor/main.go


build-all: build-user

build: init test build-all