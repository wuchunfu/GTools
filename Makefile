#  本地运行
.PHONY: dev
dev:
	wails dev

# 根据系统进行本地构建， -tags exp_gowebview2loader 需要wails v2.2.0, 低于此版本请继续使用 -webview2 browser
.PHONY: build
build:
	wails build -clean -tags exp_gowebview2loader

# 运行docsify文档服务
.PHONY: doc
doc:
	docsify serve docs --port=3010