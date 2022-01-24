# Makefile

TELEGRAM_TOKEN := $(TELEGRAM_TOKEN)
URL := "https://7260-206-190-232-194.ngrok.io"
GHCR_URL := ghcr.io
APP_NAME := pro-bot
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

.PHONY: build
build:
	@mkdir -p functions
	@go get ./...
	@go build -o functions/$(APP_NAME) .

.PHONY: help
help:
	$(info ${HELP_MESSAGE})
	@exit 0

define HELP_MESSAGE

Usage: $ make [TARGETS]

TARGETS

	help            Show the help menu
	build           Build the application artifact
	run             Run the application container locally (VERSION optional)
	set-webhook     Update the bound webhook url

EXAMPLE USAGE

	build           Build the application image and tag it as latest
	run             Run the application container locally with the latest tag
	publish         Build the application iamge, tag it as v1.0.1, and push it to GHCR

endef
