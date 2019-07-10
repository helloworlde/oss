package util

import (
	"oss/config"
	"testing"
)

func TestValidateConfig(t *testing.T) {
	rightConfig := config.OssConfig{
		EndPoint:       "EndPoint",
		AccessKeyId:    "AccessKeyId",
		AccessSecretId: "AccessSecretId",
		PicturePath:    "PicturePath",
		BucketName:     "BucketName",
		Host:           "Host",
	}

	wrongConfig := config.OssConfig{}

	result := ValidateConfig(wrongConfig)
	if result {
		t.Error("配置是不正确的")
	}

	result = ValidateConfig(rightConfig)
	if !result {
		t.Error("配置是正确的")
	}

}
