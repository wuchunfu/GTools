package gtools

import (
	"bufio"
	"changeme/util"
	"io/fs"
	"os"
)

// 获取md文档的全部内容
func (a *App) GetMdContent(path string) *util.Resp {
	fi, err := os.Stat(path)
	if err != nil {
		a.Log.Error(err.Error())
		return util.Error(err.Error())
	}
	if fi.IsDir() {
		a.Log.Warn("参数异常-md文档路径为文件夹")
		return util.Error("参数异常-md文档路径为文件夹")
	}
	b, err := os.ReadFile(path)
	if err != nil {
		a.Log.Error(err.Error())
		return util.Error(err.Error())
	} else {
		return util.Success(string(b))
	}
}

// 添加md文件夹列表
func (a *App) AddDirPath(path string) string {
	return util.AddDirPathToJdb(path)
}

// 获取md文件夹列表数据
func (a *App) GetDirList() *util.Resp {
	// 从指定的数据json中取出所有数据
	b, content := util.GetJdbContent("/path.json")
	if b {
		return util.Success(string(content))
	} else {
		a.Log.Error("path.json 数据获取失败")
		return util.Error("路径数据获取失败")
	}
}

// 保存md文件内容
func (a *App) SaveMdContent(path string, content string) *util.Resp {
	fi, err := os.Stat(path)
	if err != nil {
		return util.Error(err.Error())
	}
	if fi.IsDir() {
		return util.Error("内容写入失败")
	}
	err = os.WriteFile(path, []byte(content), fs.FileMode(os.O_RDWR|os.O_TRUNC))
	if err != nil {
		return util.Error(err.Error())
	} else {
		return util.Success(path)
	}
}

// 删除md文件夹列表
func (a *App) DelMdDir(path string) string {
	return util.DelDirPathFromJdb(path)
}

// 新建md文档
func (a *App) NewMd(path string, content string) *util.Resp {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		a.Log.Error("新建md文档异常: " + err.Error())
		return util.Error("新建md文档异常: " + err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return util.Success("OK")
}
