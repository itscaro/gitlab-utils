IMAGE=itscaro/gitlab-labeler

.PHONY: release
release: TAG ?= $(shell git describe --tags)
release: VERSION = $(or $(TAG),$(TAG),latest-dev)
release: GIT_COMMIT = $(shell git rev-list -1 HEAD)
release:
	@echo "Releasing $(IMAGE):$(VERSION) (GIT_COMMIT: $(GIT_COMMIT))"
	@docker build --pull --quiet \
		--build-arg VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(GIT_COMMIT)\
		--build-arg http_proxy=$$http_proxy \
		--build-arg https_proxy=$$https_proxy \
		-f Dockerfile \
		-t $(IMAGE):$(VERSION) . && \
	docker push $(IMAGE):$(VERSION)
