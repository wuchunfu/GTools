package gtools

import (
	"fmt"
	"gtools/configs"
	"gtools/util"
	"os"
	"strconv"
	"time"

	"golang.design/x/clipboard"
)

func (a *App) UploadScreenshot() *util.Resp {
	imgBedType := a.ConfigMap["imgbed"]["configType"]
	if imgBedType == "" {
		return util.Error(configs.NoImgBedConfigErr)
	}

	// 获取系统粘贴板中的文件地址
	b, imgPath := SavePngFromClipboard()
	if !b {
		// 删除图片临时文件
		os.Remove(imgPath)

		a.Log.Error("系统剪贴板图片地址异常: " + imgPath)
		return util.Error(configs.GetSystemClipboardImgErr)
	}

	// 判断文件是否存在并为非文件夹
	b, s := util.CheckFileExist(imgPath)
	if !b {
		a.Log.Error("系统剪贴板图片地址异常: " + s)
		return util.Error(s)
	}

	// 选择图床存储文件
	switch imgBedType {
	case "alioss":
		if b, s = util.UploadByAliOss(imgPath, a.AliOSS, a.Log, a.ConfigMap["alioss"]); !b {
			return util.Error("图片上传阿里云OSS失败, 请检查软件配置和网路状况")
		}
	case "limgpath":
		if b, s = util.MoveImgToPath(imgPath, a.ConfigMap["limgpath"]["path"]); !b {
			return util.Error("图片本地存储失败, 请检查软件配置")
		}
	default:
	}
	// 删除图片临时文件
	os.Remove(imgPath)
	if b {
		// TODO 图片地址存储到sqlite数据库中

		mdImg := "![](%v)" //  可以输出markdown原始图片
		return util.Success(fmt.Sprintf(mdImg, s))
	} else {
		return util.Error(s)
	}
}

func SavePngFromClipboard() (bool, string) {
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
