deps:
	go mod download
fmt:
	go fmt github.com/swarupdonepudi/gitr/cmd
	go fmt github.com/swarupdonepudi/gitr/lib
build: deps fmt
	go build -o bin/gitr-darwin main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/gitr-linux main.go
checksum: build
	openssl dgst -sha256 bin/gitr-darwin
	openssl dgst -sha256 bin/gitr-linux
test:
	go test -v ./...
