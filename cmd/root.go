/*
Copyright © 2019 HelloWoodes <hellowoodes@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
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

var ossConfig = OssConfig{}

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
		ossConfig = OssConfig{
			AccessKeyId:    viper.GetString("aliyun.accessKeyId"),
			AccessSecretId: viper.GetString("aliyun.accessSecretId"),
			EndPoint:       viper.GetString("aliyun.oss.endPoint"),
			Host:           viper.GetString("aliyun.oss.host"),
			PicturePath:    viper.GetString("aliyun.oss.picturePath"),
			BucketName:     viper.GetString("aliyun.oss.bucketName"),
		}
	}
}
