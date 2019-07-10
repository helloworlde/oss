package util

import (
	"testing"
)

func Test_getFileSize(t *testing.T) {
	size := int64(90225543)
	fileSize := GetFileSize(size)
	if "86.05M" != fileSize {
		t.Error("90225543 Byte 的大小应该为 86.05M")
	}

	size = int64(1610612736)
	fileSize = GetFileSize(size)
	if "1.50G" != fileSize {
		t.Error("1610612736 Byte 的大小应该为 1.50G")
	}
}

func Test_isFileExist(t *testing.T) {
	filePath := "./file_util_test.go"
	if !IsFileExist(filePath) {
		t.Error(filePath, "文件应该存在")
	}

	filePath = "notExistFile"
	if IsFileExist(filePath) {
		t.Error(filePath, "文件不应该存在")
	}
}

func Test_isDir(t *testing.T) {
	filePath := "./file_util_test.go"
	if IsDir(filePath) {
		t.Fatal(filePath, "不是文件夹")
	}

	filePath = "./"
	if !IsDir(filePath) {
		t.Error(filePath, "是文件夹")
	}

	filePath = "./tmp/"
	if IsDir(filePath) {
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
	filePath := "file_util_test.go"
	fileName := GetFileName(filePath)
	expectedFileName := "file_util_test.go"
	if fileName != expectedFileName {
		t.Errorf("文件名应该为 %s", expectedFileName)
	}

	filePath = "../main.go"
	fileName = GetFileName(filePath)
	expectedFileName = "main.go"
	if fileName != expectedFileName {
		t.Errorf("文件名应该为 %s", expectedFileName)
	}
}

func TestGetDirFiles(t *testing.T) {
	path := "../../oss"
	files := GetDirFiles(path)

	for _, v := range files {
		t.Log(v)
	}

}
