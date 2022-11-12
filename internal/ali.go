package internal

import (
	"fmt"
	"gtools/configs"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
)

func NewOssClient(endPoint string, accessKeyId string, accessKeySecret string, log *logrus.Logger) *oss.Client {
	// 创建OSSClient实例。
	client, err := oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		log.Error(fmt.Sprintf(configs.CreateAliOSSClientErr, err.Error()))
	}
	return client
}
