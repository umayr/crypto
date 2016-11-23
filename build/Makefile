GO ?= go
GOENV := $(shell $(GO) env GOOS)_$(shell $(GO) env GOARCH)
GOROOTINFO := $(shell $(GO) version | cut -d' ' -f 3)_$(GOENV)
GO_COMMIT := 4f7363cd723d3c79081d3a201279f4e024620db9
GO_BRANCH := release-branch.go1.7
SUB_PACKAGES := aes

.PHONY: GOROOT
GOROOT: GOROOT/$(GOENV)/.ok_$(GOROOTINFO)
	@for pkg in $(SUB_PACKAGES); do \
		rm -f GOROOT/$(GOENV)/pkg/*/crypto/$$pkg.a; \
	done

GOROOT/$(GOENV)/.ok_$(GOROOTINFO):
	rm -rf GOROOT/$(GOENV)
	mkdir -p GOROOT/$(GOENV)/pkg
	cp -r "$(shell $(GO) env GOROOT)/src" GOROOT/$(GOENV)/src
	cp -r "$(shell $(GO) env GOROOT)/pkg/include" GOROOT/$(GOENV)/pkg/include
	cp -r "$(shell $(GO) env GOROOT)/pkg/tool" GOROOT/$(GOENV)/pkg/tool
	
	@for pkg in $(SUB_PACKAGES); do \
		rm -r GOROOT/$(GOENV)/src/crypto/$$pkg; \
		ln -s ../../../../../$$pkg GOROOT/linux_amd64/src/crypto/; \
	done;

	GOROOT="$(CURDIR)/GOROOT/$(GOENV)" $(GO) install -v std
ifeq ($(shell go env CGO_ENABLED),1)
	GOROOT="$(CURDIR)/GOROOT/$(GOENV)" $(GO) install -race -v std
endif
	@touch "$@"

.PHONY: go
go: go/.ok_$(GO_COMMIT)_$(GOENV)

go/.ok_$(GO_COMMIT)_$(GOENV):
	rm -rf go/.ok_*_$(GOENV) go/$(GOENV)
	mkdir -p go
	git clone --branch $(GO_BRANCH) --single-branch --depth 25 https://github.com/umayr/go go/$(GOENV)
	cd go/$(GOENV) && git checkout $(GO_COMMIT)
	cd go/$(GOENV)/src && GOROOT_BOOTSTRAP="$(shell $(GO) env GOROOT)" ./make.bash
	@touch "$@"

.PHONY: clean
clean:
	rm -rf GOROOT go