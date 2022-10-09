package utils

import (
	"fmt"
	"os"
	"os/user"
)

func InitJdb(fileName string) (bool, string, error) {
	// 获取Mac Home路径
	user, _ := user.Current()
	var homeDir = user.HomeDir
	// 判断 homeDir/gtools 文件夹是否存在
	var gtDir = homeDir + "/GTools"
	_, err := os.Stat(gtDir)
	if err != nil {
		if os.IsNotExist(err) {
			err2 := os.Mkdir(gtDir, 0755)
			if err2 != nil {
				fmt.Println("文件夹创建失败！")
				return false, "", err2
			}
		} else {
			return false, "", err
		}
	}

	// 判断文件是否存在
	// 配置文件路径
	var confFile = gtDir + "/" + fileName
	_, err3 := os.Stat(confFile)
	if err3 != nil {
		if os.IsNotExist(err3) {
			_, err4 := os.Create(confFile)
			if err4 != nil {
				return false, "", err4
			} else {
				return true, confFile, err4
			}
		} else {
			return false, "", err3
		}
	}
	return true, confFile, err3
}

func GetJdbContent(jdbName string) (bool, []byte) {
	user, _ := user.Current()
	var homeDir = user.HomeDir
	var jdbPath = homeDir + "/GTools" + jdbName
	b, err := os.ReadFile(jdbPath)
	if err != nil {
		b2, _, err2 := InitJdb(jdbName)
		if b2 {
			b, _ := os.ReadFile(jdbPath)
			return true, b
		} else {
			return false, []byte(err2.Error())
		}
	} else {
		return true, b
	}
}
