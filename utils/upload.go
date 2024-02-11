package utils

import (
	"blog-go/config"
	"mime/multipart"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadFile(file *multipart.FileHeader) (string, int) {
	aliyunOSSConfig := config.GetAliyunOSSConfig()
	ossClient, err := oss.New(aliyunOSSConfig.AliyunServer, aliyunOSSConfig.AccessKey, aliyunOSSConfig.SecretKey)
	if err != nil {
		return "", UnknownErr
	}

	bucket, err := ossClient.Bucket(config.GetAliyunOSSConfig().Bucket)
	if err != nil {
		return "", UnknownErr
	}

	f, err := file.Open()
	if err != nil {
		return "", UnknownErr
	}
	defer func(f multipart.File) {
		_ = f.Close()
	}(f)

	err = bucket.PutObject(file.Filename, f)
	if err != nil {
		return "", UnknownErr
	}

	return aliyunOSSConfig.AliyunServer + "/" + file.Filename, Success
}
