package internal

import (
	"gtools/configs"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func NewOssClient() *oss.Client {
	// 创建OSSClient实例。
	client, err := oss.New(configs.Endpoint, configs.AccessKeyId, configs.AccessKeySecret)
	if err != nil {
		panic(err.Error())
	}
	return client
}
