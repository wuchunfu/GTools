# 如何使用Wails创建项目

## 安装Wails

> 官网教程：[https://wails.io/zh-Hans/docs/gettingstarted/installation](https://wails.io/zh-Hans/docs/gettingstarted/installation)

核心依赖

- Go 1.18+
- NPM (Node 15+)

### 安装Wails Cli

```shell
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### 系统检测

```shell
wails doctor
```

### 问题及异常

系统检测提示缺失wails命令，因为go install 将生成的可执行文件放在了GOPATH目录下的bin文件夹中，该文件夹没有添加到系统环境变量中，而配置golang开发环境时一般只将GOROOT添加到了系统环境变量中，所以解决办法就是：将GOPATH/bin/wails文件移动到GOROOT/bin/目录下

查看golang配置文件
重点查看 `GOPATH` 和 `GOROOT` 路径

```shell
go env
```

## 创建项目

使用指令创建

```shell
wails init
```

使用现有模板创建

```shell
wails init -n myproject -t vue
```

运行项目

```shell
wails dev
```

构建项目

```shell
wails build
```

## 更换前端

使用vue cli创建 vue2 ts 前端

1. 删除根目录下的 `frontend`文件夹
2. 使用 `vue create frontend `指令创建前端项目
3. 修改前端运行指令
4. 前端增加热更新配置
