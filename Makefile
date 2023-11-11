PACKAGE = github.com/xuender/grephub
VERSION = $(shell git describe --tags)
BUILD_TIME = $(shell date +%F' '%T)

default: lint-fix test

tools:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/cespare/reflex@latest

lint:
	golangci-lint run --timeout 60s --max-same-issues 50 ./...

lint-fix:
	golangci-lint run --timeout 60s --max-same-issues 50 --fix ./...

test:
	go test -race -v ./... -gcflags=all=-l -cover

watch-test:
	reflex -t 50ms -s -- sh -c 'go test -race -v ./...'

clean:
	rm -rf dist
	rm -rf app/www

proto: protojs
	protoc --go_out=. pb/*.proto

protojs:
	cd frontend && node_modules/.bin/pbjs -t static-module -w commonjs -o src/pb.js ../pb/*.proto
	cd frontend && node_modules/.bin/pbts -o src/pb.d.ts src/pb.js

build: clean
	./frontend/node_modules/.bin/ng build --base-href ./

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}-linux-amd64 cmd/grephub/main.go

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}-darwin-amd64 cmd/grephub/main.go

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
	go build \
	-ldflags "-s -w -H=windowsgui -X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}.exe cmd/grephub/main.go

b:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go build \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}-linux-amd64 cmd/grephub/main.go

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 \
	go build \
	-ldflags "-X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}-darwin-amd64 cmd/grephub/main.go

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 \
	go build \
	-ldflags "-s -w -H=windowsgui -X 'github.com/xuender/kit/oss.Version=${VERSION}' \
  -X 'github.com/xuender/kit/oss.BuildTime=${BUILD_TIME}'" \
  -o dist/grephub-${VERSION}.exe cmd/grephub/main.go

wire:
	wire gen ${PACKAGE}/app