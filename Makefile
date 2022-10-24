#  本地运行
.PHONY: dev
dev:
	wails dev

# Windows本地构建
.PHONY: windows
buildw:
	wails build -clean -platform windows/amd64 -o gtools.exe -webview2 browser -tags "CGO_ENABLE=1 GOOS=windows GOARCH=amd64"

# Mac本地构建
.PHONY: mac-int
buildd:
	wails build -clean -platform darwin/amd64

# 运行docsify文档服务
.PHONY: doc
doc:
	docsify serve docs --port=3010