package cmd

type OssConfig struct {
	EndPoint       string `valid:"required"`
	AccessKeyId    string `valid:"required"`
	AccessSecretId string `valid:"required"`
	BucketName     string `valid:"required"`
	Host           string `valid:"required"`
	PicturePath    string `valid:"required"`
}
