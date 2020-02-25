compile:
	go mod download
	go build -o bin/gitr-darwin main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/gitr-linux main.go
