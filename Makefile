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
setup-tests:
	mv lib_test/test_data/r1-no-remote/.git-temp lib_test/test_data/r1-no-remote/.git
	mv lib_test/test_data/r2-with-remote/.git-temp lib_test/test_data/r2-with-remote/.git
	mv lib_test/test_data/r3-with-remote-custom-branch/.git-temp lib_test/test_data/r3-with-remote-custom-branch/.git
execute-tests:
	go test -v github.com/swarupdonepudi/gitr/lib_test || true
cleanup-tests:
	mv lib_test/test_data/r1-no-remote/.git lib_test/test_data/r1-no-remote/.git-temp
	mv lib_test/test_data/r2-with-remote/.git lib_test/test_data/r2-with-remote/.git-temp
	mv lib_test/test_data/r3-with-remote-custom-branch/.git lib_test/test_data/r3-with-remote-custom-branch/.git-temp
test: setup-tests execute-tests cleanup-tests
