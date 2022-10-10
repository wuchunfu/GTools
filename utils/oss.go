package utils

import (
	"fmt"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadImgFromClipboard() string {
	var resp Resp[string]
	resp.Code = 500

	// 获取系统粘贴板中的文件地址
	b, imgPath := SavePngFromClipboard()
	if !b {
		resp.Msg = imgPath
		return resp.toJson()
	}

	fmt.Printf("imgPath: %v\n", imgPath)

	// 判断文件是否存在并为非文件夹
	b, s := CheckFileExist(imgPath)
	if !b {
		resp.Msg = s
		return resp.toJson()
	}

	// 从配置文件中选取上传首选项 TODO 目前先写死只使用阿里云OSS上传图片
	b2, s2 := uploadByAliOss(imgPath)
	if b2 {
		resp.Code = 200
		resp.Data = s2
		return resp.toJson()
	} else {
		resp.Msg = s2
		return resp.toJson()
	}
}

func uploadByAliOss(path string) (bool, string) {
	// 获取文件名称
	s := strings.Split(path, "/")
	var fileName = s[len(s)-1]

	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	confEndpoint := "oss-cn-beijing.aliyuncs.com"
	endpoint := "https://" + confEndpoint

	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "LTAI5tF8bKmV9Be9axkMgM5b"
	accessKeySecret := "FrCuS7gqW1HW1zYVe3iKdnTLuEnmym"

	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		return false, err.Error()
	}

	// yourBucketName填写存储空间名称。
	bucketName := "ly-img-xiao"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。 注意test前不能有 '/' 符号
	// 最后获得的图片地址为：endpoint + "/" + objectName
	objectName := "test/" + fileName
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := path

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return false, err.Error()
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		return false, err.Error()
	}

	return true, "https://"+ bucketName + "." + confEndpoint + "/" + objectName
}
