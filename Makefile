init:
	go mod download

generate:
	go generate ./...

test:
	go test -v ./... -cover

build-user:
	mkdir -p bin/ && go build -o bin/user app/user/cmd/main.go

build-netlify: build-user

publish:
	mkdir -p swagger/ && cp api/user/v1/user.swagger.yaml swagger/

build: init test build-netlify publish