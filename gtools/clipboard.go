package gtools

import (
	"changeme/util"
	"fmt"
	"os"
	"strconv"
	"time"

	"golang.design/x/clipboard"
)

func (a *App) UploadScreenshot() *util.Resp {

	// 获取系统粘贴板中的文件地址
	b, imgPath := savePngFromClipboard()
	if !b {
		a.Log.Error("系统剪贴板图片地址异常: " + imgPath)
		return util.Error(imgPath)
	}

	// 判断文件是否存在并为非文件夹
	b, s := util.CheckFileExist(imgPath)
	if !b {
		a.Log.Error("系统剪贴板图片地址异常: " + s)
		return util.Error(s)
	}

	// TODO 从配置文件中选取上传首选项  目前先写死只使用阿里云OSS上传图片
	b, s = util.UploadByAliOss(imgPath, a.AliOSS, a.Log)
	if b {
		// 删除图片临时文件
		os.Remove(imgPath)
		// TODO 将线上图片地址存储到sqlite数据库中

		mdImg := "![](%v)" //  可以输出markdown原始图片
		return util.Success(fmt.Sprintf(mdImg, s))
	} else {
		return util.Error(s)
	}
}

func savePngFromClipboard() (bool, string) {
	b2 := clipboard.Read(clipboard.FmtImage)
	i := len(b2)
	if i != 0 {
		unix := strconv.FormatInt(time.Now().Local().UnixMilli(), 10)
		imgPath := fmt.Sprintf("%s/%s.png", util.GetConfigDir(), unix)
		f, err := os.Create(imgPath)
		f.Write(b2)
		defer f.Close()
		if err != nil {
			return false, err.Error()
		}
		return true, imgPath
	} else {
		return false, "系统粘贴板中没有图片！"
	}
}
