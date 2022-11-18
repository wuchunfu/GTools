package util

import (
	"fmt"
	"gtools/configs"
	"os"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

func UploadByAliOss(path string, client *oss.Client, log *logrus.Logger, config map[string]string) (bool, string) {
	// 获取文件名称
	s := strings.Split(path, "/")
	var fileName = s[len(s)-1]

	// 生成Bucket下的图片路径
	objectName := fmt.Sprintf("%s/%s", config["projectDir"], fileName)

	// 获取存储空间。
	bucket, err := client.Bucket(config["bucketName"])
	if err != nil {
		log.Error(fmt.Sprintf(configs.GetAliOSSBucketErr, err.Error()))
		return false, err.Error()
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, path)
	if err != nil {
		log.Error("AliOSS图片上传失败: " + err.Error())
		return false, err.Error()
	}
	return true, fmt.Sprintf("https://%s.%s/%s", config["bucketName"], config["endPoint"], objectName)
}

// 将临时图片文件存储到指定的文件夹下
func MoveImgToPath(imgPath string, targetPath string) (bool, string) {
	if targetPath == "" {
		return false, "本地存储路径未配置"
	}
	s := strings.Split(imgPath, "/")
	imgName := s[len(s)-1]
	newPath := fmt.Sprintf("%s/%s", targetPath, imgName)
	if err := os.Rename(imgPath, newPath); err != nil {
		return false, configs.SaveImgToLocalErr
	}
	// return true, imgName
	return true, newPath
}
