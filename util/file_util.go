package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 判断文件是否存在
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断是否是文件夹
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// 获取绝对路径
func getAbsolutePath(path string) string {
	dir, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("获取%s绝对路径失败", path)
	}

	return strings.Replace(dir, "\\", "/", -1)
}

// 获取文件名
func GetFileName(path string) string {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return ""
	}
	return fileInfo.Name()
}

const K = 1024
const M = 1024 * 1024
const G = M * 1024

// 获取文件大小
func GetFileSize(size int64) interface{} {

	if size > G {
		return fmt.Sprintf("%.2f", float64(size)/float64(G)) + "G"
	} else if size > M {
		return fmt.Sprintf("%.2f", float64(size)/float64(M)) + "M"
	}
	return fmt.Sprintf("%.2f", float64(size)/float64(K)) + "K"
}

func GetDirFiles(dir string) []string {

	if !IsDir(dir) {
		return nil
	}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var fileArray = make([]string, len(files))

	for _, file := range files {
		if file.IsDir() {
			subDir := GetDirFiles(dir + "/" + file.Name())
			if subDir != nil {
				fileArray = append(fileArray, subDir...)
			}
		} else {
			fileArray = append(fileArray, dir+"/"+file.Name())
		}

	}

	return fileArray
}
