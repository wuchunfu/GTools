# GTools开发说明

## 简介

GTools是基于wails开发的一款跨平台（Mac、Windows、Linux）桌面小工具

## 起源

一款目前来说只适合我自己的跨平台小工具。

## 开发

- MacOS/Windows
- Vscode
- wails v2.1.0
- golang 1.18
- Vue 3.2.37
- Naive UI
- Element-plus

## 文件结构

```
|-- GTools
    |-- build               # 构建文件夹，生成可执行文件
        |-- appicon.png     # 软件图标
    |-- configs             # 存放静态文件或者静态变量
    |-- docs                # docsify文档目录，存放使用说明和开发说明文档
    |-- frontend            # vue前端工程文件夹，可单独使用WS进行开发
    |-- gtools              # 存放前端需要的接口，类似controller（个人喜好）
    |-- internal            # 存放一些基本配置文件（xorm、log...）
    |-- test                # 存放一些测试文件
    |-- util                # 存放一些工具类
    |-- main.go             # 程序入口
    |-- Makefile            # 构建指令
    |-- wails.json          # wails配置文件
```

## 目前功能点

1. Markdown编辑器 - 主要由Vue3实现，使用了功能较为齐全的Vditor
2. 快捷指令 - 主要存储并执行常用的快捷指令，基于vue-web-terminal
3. 设置 - 主要进行软件的功能设置

## TODO

- [x] Markdown编辑器增加阿里云OSS图片上传功能
- [x] Markdown编辑器增加目录功能
- [x] Markdown编辑器增加失去焦点保存功能
- [x] Markdown编辑器增加创建、保存文档功能
- [x] 使用Naive UI下拉框实现Markdown编辑器的文件树功能
- [x] 实现添加文件夹和文件功能
- [ ] Markdown文本的PDF、HTML导出功能
- [ ] markdown基于百度翻译的实时翻译功能
- [x] 将markdown的存储方式由json改为sqlite3
- [x] 重构Golang后端代码结构
- [ ] 实现大部分设置的动态配置
- [ ] 增加Markdown编辑器图片上传和本地路径存储的选项
- [ ] 富文本编辑器的文件增删改查功能实现
- [x] 使用xorm替代json文件进行数据的存储
- [ ] 增加时钟功能
- [ ] 增加倒计时结束后系统通知功能
- [ ] 使用go-toast实现系统通知功能
- [ ] 多线程文件下载器
- [ ] 增加快捷笔记功能
- [ ] 增加PDF阅读器功能(非必要)
- [ ] 快捷指令增、删、改、查、运行、停止功能
- [x] 待办事项清单分组、增、删、改、查功能
- [ ] Json格式化小工具
- [ ] 图片-Base64转换工具
- [ ] 房贷计算器小工具
- [ ] 基于百度OCR的图片精准转文字小工具
- [ ] 更多的功能拓展
- [ ] 精简前端后端包结构，继续减小包体积
