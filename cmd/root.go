package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"oss/config"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "oss",
	Short: "For upload file to Aliyun OSS in command line.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var ossConfig = config.OssConfig{}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.oss.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".oss" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".oss")
	}

	viper.AutomaticEnv() // read in environment variables that match

	if err := viper.ReadInConfig(); err == nil {
		// 初始化配置
		ossConfig = config.OssConfig{
			AccessKeyId:    viper.GetString("aliyun.accessKeyId"),
			AccessSecretId: viper.GetString("aliyun.accessSecretId"),
			EndPoint:       viper.GetString("aliyun.oss.endPoint"),
			Host:           viper.GetString("aliyun.oss.host"),
			PicturePath:    viper.GetString("aliyun.oss.picturePath"),
			BucketName:     viper.GetString("aliyun.oss.bucketName"),
		}
	}
}
