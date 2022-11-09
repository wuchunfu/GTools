package util

import (
	"fmt"
	"gtools/configs"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

func UploadByAliOss(path string, client *oss.Client, log *logrus.Logger) (bool, string) {
	// 获取文件名称
	s := strings.Split(path, "/")
	var fileName = s[len(s)-1]

	// 生成Bucket下的图片路径
	objectName := fmt.Sprintf("%s/%s", configs.ProjectDir, fileName)

	// 获取存储空间。
	bucket, err := client.Bucket(configs.BucketName)
	if err != nil {
		log.Error("获取AliOSS存储空间失败: " + err.Error())
		return false, err.Error()
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, path)
	if err != nil {
		log.Error("AliOSS图片上传失败: " + err.Error())
		return false, err.Error()
	}
	return true, fmt.Sprintf("https://%s.%s/%s", configs.BucketName, configs.Endpoint, objectName)
}
