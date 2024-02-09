package oss

import (
	"blog-go/config"
	"blog-go/pkg/utils"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var ossClient *oss.Client

func UploadFile(file *multipart.FileHeader) (string, int) {
	aliyunOSSConfig := config.GetAliyunOSSConfig()
	ossClient, err := oss.New(aliyunOSSConfig.AliyunServer, aliyunOSSConfig.AccessKey, aliyunOSSConfig.SecretKey)
	if err != nil {
		return "", utils.UnknownErr
	}

	bucket, err := ossClient.Bucket(config.GetAliyunOSSConfig().Bucket)
	if err != nil {
		return "", utils.UnknownErr
	}

	f, err := file.Open()
	if err != nil {
		return "", utils.UnknownErr
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f)

	err = bucket.PutObject(file.Filename, f)
	if err != nil {
		return "", utils.UnknownErr
	}

	return aliyunOSSConfig.AliyunServer + "/" + file.Filename, utils.Success
}
