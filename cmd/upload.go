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
	"errors"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload [file]",
	Short: "Upload file to OSS",
	Long:  ``,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one file path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir := cmd.Flag("directory").Value.String()
		picture := cmd.Flag("picture").Value.String()

		ossDir := ""

		if dir != "" {
			ossDir = dir
		} else if picture == "true" {
			ossDir = ossConfig.PicturePath
		}

		var files = [10]string{}

		length := len(args)
		if length > 0 {

			for index, file := range args {
				filePath := string(file)
				if !isFileExist(string(filePath)) {
					fmt.Println("文件", filePath, "不存在")
					os.Exit(1)
				}
				if isDir(filePath) {
					//TODO 文件夹上传
					fmt.Println("暂不支持上传文件夹")
					os.Exit(1)
				} else {
					files[index] = filePath
				}
			}
		}

		validateConfig()

		bucket := initBucket(ossConfig)

		for _, path := range files {
			if len(path) > 0 {
				UploadFile(ossDir, path, ossConfig, *bucket)
			}
		}
	},
}

func validateConfig() {

	result, err := govalidator.ValidateStruct(ossConfig)
	if !result || err != nil {
		errMessage := strings.ReplaceAll(err.Error(), ";", "\n ")
		errMessage = strings.ReplaceAll(errMessage, "non zero value required", "配置无效")
		fmt.Println("Config not correct:\n", errMessage)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("directory", "d", "", "Dictionary of upload file")
	uploadCmd.Flags().BoolP("picture", "p", true, "Upload to Picture folder")
}
