.PHONY: dev
dev:
	wails dev

.PHONY: build
build:
	wails build -clean -platform darwin,windows

.PHONY: doc
doc:
	docsify serve docs --port=3010