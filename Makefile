name=gitr
name_local=gitr
pkg=github.com/plantoncloud/gitr
build_dir=build
v?=dev
LDFLAGS=-ldflags "-X ${pkg}/internal/version.Version=${v}"
build_cmd=go build -v ${LDFLAGS}

.PHONY: deps
deps:
	go mod download

.PHONY: build
build: ${build_dir}/${name}

${build_dir}/${name}: deps
	GOOS=darwin ${build_cmd} -o ${build_dir}/${name}-darwin .
	GOOS=darwin GOARCH=amd64 ${build_cmd} -o ${build_dir}/${name}-darwin-amd64 .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-amd64
	GOOS=darwin GOARCH=arm64 ${build_cmd} -o ${build_dir}/${name}-darwin-arm64 .
	openssl dgst -sha256 ${build_dir}/${name}-darwin-arm64
.PHONY: test
test:
	go test -race -v -count=1 ./...

.PHONY: run
run: build
	${build_dir}/${name}

.PHONY: vet
vet:
	go vet ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: clean
clean:
	rm -rf ${build_dir}

checksum:
	@openssl dgst -sha256 ${build_dir}/${name}-darwin

local:
	sudo rm -f /usr/local/bin/${name_local}
	sudo cp ./${build_dir}/${name}-darwin /usr/local/bin/${name_local}
	sudo chmod +x /usr/local/bin/${name_local}

release: build
	gsutil -h "Cache-Control:no-cache" cp build/gitr-darwin-amd64 gs://planton-pcs-artifact-file-repo/tool/gitr/download/gitr-${v}-amd64
	gsutil -h "Cache-Control:no-cache" cp build/gitr-darwin-arm64 gs://planton-pcs-artifact-file-repo/tool/gitr/download/gitr-${v}-arm64
