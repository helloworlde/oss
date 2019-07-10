package cmd

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"strconv"
	"strings"
)

// 上传文件
func UploadFile(ossPath string, localFilePath string, config OssConfig, bucket oss.Bucket) (string, string) {
	name := getFileName(localFilePath)

	// 移除/替换多余的 /
	objectName := strings.ReplaceAll(ossPath+"/"+name, "//", "/")
	objectName = strings.TrimPrefix(objectName, "/")

	err := bucket.PutObjectFromFile(objectName, localFilePath, oss.Progress(&OssProgressListener{}))

	if err != nil {
		fmt.Println("上传文件失败:", err)
		os.Exit(1)
	}

	ossHost := config.Host

	return localFilePath, ossHost + "/" + objectName
}

// 初始化 Bucket
func initBucket(config OssConfig) *oss.Bucket {
	endPoint := config.EndPoint
	accessKeyId := config.AccessKeyId
	accessSecretId := config.AccessSecretId
	bucketName := config.BucketName

	client, err := oss.New(endPoint, accessKeyId, accessSecretId)
	if err != nil || client == nil {
		fmt.Println("初始化 OSS Client 失败:", err)
		os.Exit(1)
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil || bucket == nil {
		fmt.Println("初始化 OSS Bucket 失败:", err)
		os.Exit(1)
	}
	return bucket
}

// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferDataEvent:
		rate := strconv.FormatInt(event.ConsumedBytes*100/event.TotalBytes, 10) + "%"
		fmt.Printf("\rTransfer Rate:%s TotalSize: %s FinishedSize: %-8s", rate, getFileSize(event.TotalBytes), getFileSize(event.ConsumedBytes))
	default:
	}
}
