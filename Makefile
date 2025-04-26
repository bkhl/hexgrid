MAKEFLAGS += --always-make

RUNNER := podman

GO_CACHE_VOLUME := hexgrid-build

GO_LINT_IMAGE := docker.io/golangci/golangci-lint:latest

GO_IMAGE := docker.io/library/golang:1.24.2

define run
$(RUNNER) run --rm --interactive \
--env GOCACHE=/var/lib/go/cache \
--env GOPATH=/var/lib/go \
--volume $(GO_CACHE_VOLUME):/var/lib/go:z \
--volume $(CURDIR):$(CURDIR):z \
--workdir $(CURDIR)
endef

help:
	-make --print-targets

go:
	$(run) $(GO_IMAGE) go $(ARGS)

lint:
	$(run) $(GO_LINT_IMAGE) golangci-lint run -v

test:
	$(run) $(GO_IMAGE) go test

clean:
	podman volume rm -f $(GO_CACHE_VOLUME)
