package gtools

import (
	"bufio"
	"fmt"
	"gtools/configs"
	"gtools/internal"
	"gtools/util"
	"io/fs"
	"os"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MdInfo struct {
	Fname string `json:"fname"`
	Fpath string `json:"fpath"`
}

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
func (a *App) AddMdDirPath(item internal.MdPath) *util.Resp {
	_, err := a.Db.Insert(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetMdDirList()
}

// 获取文件夹列表
func (a *App) GetMdDirList() *util.Resp {
	mdDirList := make([]internal.MdPath, 0)
	mdFileList := make([]internal.MdPath, 0)
	err := a.Db.Where("type = ?", 0).Find(&mdDirList)
	if err != nil {
		return util.Error(err.Error())
	}
	err2 := a.Db.Where("type = ?", 1).Find(&mdFileList)
	if err2 != nil {
		return util.Error(err2.Error())
	}
	resultMap := make(map[string][]internal.MdPath, 0)
	resultMap["dir"] = mdDirList
	resultMap["file"] = mdFileList
	return util.Success(resultMap)
}

// 获取文件夹下所有的md文件
func (a *App) GetMdFileList(dirPath string) *util.Resp {
	mdFileList := make([]map[string]string, 0)
	de, err := os.ReadDir(dirPath)
	if err != nil {
		return util.Error(err.Error())
	}
	for _, v := range de {
		s := strings.Split(v.Name(), ".")
		if s[len(s)-1] == "md" {
			mapPath := make(map[string]string)
			fname := v.Name()
			mapPath["fname"] = fname
			mapPath["fpath"] = fmt.Sprintf("%s/%s", dirPath, v.Name())
			mdFileList = append(mdFileList, mapPath)
		}
	}
	return util.Success(mdFileList)
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
func (a *App) DelMdDir(item internal.MdPath) *util.Resp {
	_, err := a.Db.Delete(&item)
	if err != nil {
		return util.Error(err.Error())
	}
	return a.GetMdDirList()
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
	if path != "" {
		s := strings.Split(path, "/")
		item := internal.MdPath{
			Path:  path,
			Type:  1,
			Fname: s[len(s)-1],
		}
		_, err = a.Db.Insert(&item)
		if err != nil {
			return util.Error(err.Error())
		}
	}
	return a.GetMdDirList()
}

// 添加文件
func (a *App) AddMdFile() *util.Resp {
	options := runtime.OpenDialogOptions{
		Title:                "选择文件",
		Filters:              [](runtime.FileFilter){configs.MdFilter},
		CanCreateDirectories: true,
	}
	path, err := runtime.OpenFileDialog(a.ctx, options)
	if err != nil {
		return util.Error(err.Error())
	}
	if path != "" {
		s := strings.Split(path, "/")
		item := internal.MdPath{
			Path:  path,
			Type:  1,
			Fname: s[len(s)-1],
		}
		_, err = a.Db.Insert(&item)
		if err != nil {
			return util.Error(err.Error())
		}
	}
	return a.GetMdDirList()
}

func (a *App) ExportHTML(fpath, content string) *util.Resp {
	file, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		a.Log.Error("创建HTML文件异常: " + err.Error())
		return util.Error("创建HTML文件异常: " + err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(content)
	writer.Flush()
	return util.Ok()
}
