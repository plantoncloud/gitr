.PHONY: deps
deps:
	go mod download
.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...
.PHONY: build
build: deps vet fmt
	env GOOS=darwin GOARCH=amd64 go build -o bin/gitr-darwin main.go
	env GOOS=linux GOARCH=amd64 go build -o bin/gitr-linux main.go
	env GOOS=windows GOARCH=386  go build -o bin/gitr-windows-386.exe main.go
.PHONY: checksum
checksum: build
	openssl dgst -sha256 bin/gitr-darwin
.PHONY: setup-tests
setup-tests:
	mv internal/git_test_data/r1-no-remote/.git-temp internal/git_test_data/r1-no-remote/.git
	mv internal/git_test_data/r2-with-remote/.git-temp internal/git_test_data/r2-with-remote/.git
	mv internal/git_test_data/r3-with-remote-custom-branch/.git-temp internal/git_test_data/r3-with-remote-custom-branch/.git
.PHONY: execute-tests
execute-tests:
	go test -v -coverpkg github.com/swarupdonepudi/gitr/v2/internal  -cover ./... -coverprofile=internal.cov || true
	go test -v -coverpkg github.com/swarupdonepudi/gitr/v2/pkg/...  -cover ./... -coverprofile=pkg.cov || true
.PHONY: cleanup-tests
cleanup-tests:
	mv internal/git_test_data/r1-no-remote/.git internal/git_test_data/r1-no-remote/.git-temp
	mv internal/git_test_data/r2-with-remote/.git internal/git_test_data/r2-with-remote/.git-temp
	mv internal/git_test_data/r3-with-remote-custom-branch/.git internal/git_test_data/r3-with-remote-custom-branch/.git-temp
.PHONY: test
test: setup-tests execute-tests cleanup-tests
.PHONY: analyze-tests
analyze-tests:
	go tool cover -func=internal.cov
	go tool cover -func=pkg.cov
.PHONY: local
local: build
	cp bin/gitr-darwin /usr/local/bin/gitr
