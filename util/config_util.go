package util

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/helloworlde/oss/config"
	"strings"
)

// 校验配置是否正确
func ValidateConfig(ossConfig config.OssConfig) bool {

	result, err := govalidator.ValidateStruct(ossConfig)
	if !result || err != nil {
		errMessage := strings.ReplaceAll(err.Error(), ";", "\n ")
		errMessage = strings.ReplaceAll(errMessage, "non zero value required", "配置无效")
		fmt.Println("Config not correct:\n", errMessage)
		return false
	}
	return true
}
