# Makefile

TELEGRAM_TOKEN := $(TELEGRAM_TOKEN)
URL := "https://7260-206-190-232-194.ngrok.io"
GHCR_URL := ghcr.io
GHCR_USERNAME := TechProber
VERSION := latest

ifneq ($(VERSION), latest)
	export IMAGE_TAG=$(VERSION)
else
	export IMAGE_TAG=latest
endif

.PHONY: set-webhook
set-webhook:
	@echo "==> Update webhook endpoint"
	@curl -F "url=$(URL)" https://api.telegram.org/bot$(TELEGRAM_TOKEN)/setWebhook

ghcr-login:
	@echo "==> Login to GHCR"
	@echo $(GHCR_TOKEN) | docker login $(GHCR_URL) -u $(GHCR_USERNAME) --password-stdin

.PHONY: build
build:
	@echo "==> Build application image with tag $(IMAGE_TAG)"
	@DOCKER_BUILDKIT=1 docker build \
		--platform=linux/amd64 \
		-t $(IMAGE_NAME):$(IMAGE_TAG) \
		.

.PHONY: nerd-build
nerd-build:
	@echo "==> Build application image with tag $(IMAGE_TAG)"
	@sudo nerdctl build \
		--platform=linux/amd64 \
		-t $(IMAGE_NAME):$(IMAGE_TAG) \
		.

.PHONY: tag
tag:
	@echo "==> Tag the local image as GHCR image"
	@docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(GHCR_URL):$(IMAGE_TAG)
	@docker tag $(IMAGE_NAME):$(IMAGE_TAG) $(GHCR_URL):latest

.PHONY: help
help:
	$(info ${HELP_MESSAGE})
	@exit 0

define HELP_MESSAGE

Usage: $ make [TARGETS]

TARGETS

	help            Show the help menu
	build           Build the application image
	run             Run the application container locally (VERSION optional)
	set-webhook     Update the bound webhook url
	publish         Build the application image, tag it with a custom version tag, and push it to GHCR (Version required)

EXAMPLE USAGE

	build           Build the application image and tag it as latest
	run             Run the application container locally with the latest tag
	publish         Build the application iamge, tag it as v1.0.1, and push it to GHCR

endef
