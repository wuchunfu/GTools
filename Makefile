#  本地运行
.PHONY: dev
dev:
	wails dev

# 根据系统进行本地构建
.PHONY: build
build:
	wails build -clean -webview2 browser

# 运行docsify文档服务
.PHONY: doc
doc:
	docsify serve docs --port=3010