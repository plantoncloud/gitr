# ── project metadata ────────────────────────────────────────────────────────────
name        := gitr
pkg         := github.com/plantoncloud/gitr
build_dir   := build
LDFLAGS     := -ldflags "-X $(pkg)/cmd/gitr/root/version.VersionLabel=$$(git describe --tags --always --dirty)"

# ── helper vars ────────────────────────────────────────────────────────────────
build_cmd   := go build $(LDFLAGS)

# ── quality / housekeeping ─────────────────────────────────────────────────────
.PHONY: deps vet fmt test clean
deps:          ## download & tidy modules
	go mod download
	go mod tidy

vet:           ## go vet
	go vet ./...

fmt:           ## go fmt
	go fmt ./...

test: vet      ## run tests with race detector
	go test -race -v -count=1 ./...

clean:         ## remove build artifacts
	rm -rf $(build_dir)

# ── local utility ──────────────────────────────────────────────────────────────
.PHONY: snapshot local
snapshot: deps ## build a local snapshot using GoReleaser
	goreleaser release --snapshot --clean --skip-publish

local: snapshot ## copy binary to ~/bin for quick use
	install -m 0755 $(build_dir)/gitr_*_$(shell uname -m)/gitr $(HOME)/bin/$(name)

# ── release tagging ────────────────────────────────────────────────────────────
.PHONY: release build-check
build-check:   ## quick compile to verify build
	go build -o /dev/null ./cmd/$(name)

release: test build-check ## tag & push if everything passes
ifndef version
	$(error version is not set. Use: make release version=vX.Y.Z)
endif
	git tag -a $(version) -m "$(version)"
	git push origin $(version)

# ── default target ─────────────────────────────────────────────────────────────
.DEFAULT_GOAL := test
