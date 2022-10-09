# GTools开发说明

## 简介

GTools是基于wails开发的一款跨平台（Mac、Windows）桌面小工具

## 起源

本人是一名Java开发者，平时大部分的文档都是使用Markdown进行编辑，其中Typora是我最主要的编辑软件，但是由于之后的收费原因，我更换成了MarkText，但是这两款软件在使用过程中多少都有一些让我不能适应的地方，所以我决心自己开发一款真正能适合我自己的Markdown编辑器，同时也可以集成一些我常用的小功能。

## 开发环境

- MacOS
- Vscode
- wails v2.0.0-rc.1.1
- golang 1.18
- Vue 3.2.37
- Naive UI
- Element-plus


## 目前功能点

1. Markdown编辑器 - 主要由Vue3实现，使用了功能较为齐全的Vditor
2. 快捷指令 - 主要存储并执行常用的快捷指令，基于vue-web-terminal
3. 设置 - 主要进行软件的功能设置

## TODO

- [ ] Markdown编辑器增加阿里云OSS图片上传功能
- [x] Markdown编辑器增加目录功能
- [ ] Markdown编辑器增加实时保存功能
- [ ] Markdown编辑器增加创建、保存文档功能
- [ ] 快捷指令增、删、改、查功能
- [ ] 待办事项清单分组、增、删、改、查功能
- [ ] Json格式化小工具
- [ ] 图片-Base64转换工具
- [ ] 更多的功能拓展
