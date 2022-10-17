# GTools开发说明

## 简介

GTools是基于wails开发的一款跨平台（Mac、Windows）桌面小工具

## 起源

一款目前来说只适合我自己的跨平台小工具。

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

- [x] Markdown编辑器增加阿里云OSS图片上传功能
- [x] Markdown编辑器增加目录功能
- [x] Markdown编辑器增加失去焦点保存功能
- [ ] Markdown编辑器增加创建、保存文档功能
- [ ] 快捷指令增、删、改、查功能
- [ ] 待办事项清单分组、增、删、改、查功能
- [ ] Json格式化小工具
- [ ] 图片-Base64转换工具
- [ ] 更多的功能拓展
