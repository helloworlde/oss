package cmd

import (
	"testing"
)

func Test_getFileSize(t *testing.T) {
	size := int64(90225543)
	fileSize := getFileSize(size)
	if "86.05M" != fileSize {
		t.Error("90225543 Byte 的大小应该为 86.05M")
	}

	size = int64(1610612736)
	fileSize = getFileSize(size)
	if "1.50G" != fileSize {
		t.Error("1610612736 Byte 的大小应该为 1.50G")
	}
}

func Test_isFileExist(t *testing.T) {
	filePath := "./fileutil_test.go"
	if !isFileExist(filePath) {
		t.Error(filePath, "文件应该存在")
	}

	filePath = "notExistFile"
	if isFileExist(filePath) {
		t.Error(filePath, "文件不应该存在")
	}
}

func Test_isDir(t *testing.T) {
	filePath := "./fileutil_test.go"
	if isDir(filePath) {
		t.Fatal(filePath, "不是文件夹")
	}

	filePath = "./"
	if !isDir(filePath) {
		t.Error(filePath, "是文件夹")
	}

	filePath = "./tmp/"
	if isDir(filePath) {
		t.Error(filePath, "不存在")
	}
}

func Test_getAbsolutePath(t *testing.T) {
	path := "main.go"
	absPath := getAbsolutePath(path)
	if len(absPath) == 0 {
		t.Errorf("未获取到%s的绝对路径", path)
	}

}

func Test_getFileName(t *testing.T) {
	filePath := "fileutil_test.go"
	fileName := getFileName(filePath)
	expectedFileName := "fileutil_test.go"
	if fileName != expectedFileName {
		t.Errorf("文件名应该为 %s", expectedFileName)
	}

	filePath = "../main.go"
	fileName = getFileName(filePath)
	expectedFileName = "main.go"
	if fileName != expectedFileName {
		t.Errorf("文件名应该为 %s", expectedFileName)
	}
}
