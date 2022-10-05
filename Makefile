.PHONY: dev
dev:
	wails dev

.PHONY: build
build:
	wails build -clean -platform darwin/amd64,windows/amd64,linux/amd64