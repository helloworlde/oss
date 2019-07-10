package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"oss/util"
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
		markdown := cmd.Flag("markdown").Value.String()

		ossDir := ""

		if dir != "" {
			ossDir = dir
		} else if picture == "true" {
			ossDir = ossConfig.PicturePath
		}

		var files = make([]string, len(args))

		length := len(args)
		if length > 0 {

			for _, file := range args {
				filePath := string(file)
				if !util.IsFileExist(string(filePath)) {
					fmt.Println("文件", filePath, "不存在")
					os.Exit(1)
				}
				if util.IsDir(filePath) {
					dirFiles := util.GetDirFiles(filePath)
					files = append(files, dirFiles...)
				} else {
					files = append(files, filePath)
				}
			}
		}

		if !util.ValidateConfig(ossConfig) {
			os.Exit(1)
		}

		bucket := util.InitBucket(ossConfig)

		resultMap := make(map[string]string)

		for _, path := range files {
			if len(path) > 0 {
				localPath, url := util.UploadFile(ossDir, path, ossConfig, *bucket)
				resultMap[localPath] = url
			}
		}

		if markdown == "true" {
			util.PrintInMarkdownFormat(resultMap)
		} else {
			util.PrintInTextPlain(resultMap)
		}

	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)

	uploadCmd.Flags().StringP("directory", "d", "", "Dictionary of upload file")
	uploadCmd.Flags().BoolP("picture", "p", true, "Upload to Picture folder")
	uploadCmd.Flags().BoolP("markdown", "m", false, "Print link as Markdown")
}
