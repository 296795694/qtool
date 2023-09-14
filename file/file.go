package file

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// PathExists 判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CopyDir 复制整个文件夹内的文件 并替换字符串
func CopyDir(srcPath string, destPath string, fm string, to string) error {
	srcPath = strings.Replace(srcPath, "\\", "/", -1)
	destPath = strings.Replace(destPath, "\\", "/", -1)
	if srcInfo, err := os.Stat(srcPath); err != nil {
		return err
	} else {
		if !srcInfo.IsDir() {
			e := errors.New(srcPath + "不是一个正确的目录！")
			return e
		}
	}
	err := filepath.Walk(srcPath, func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if !f.IsDir() {
			path := strings.Replace(path, "\\", "/", -1)
			destNewPath := strings.Replace(path, srcPath, destPath, -1)
			CopyFile(path, destNewPath, fm, to)
		}
		return nil
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

//生成目录并拷贝文件
func CopyFile(src string, dest string, fm string, to string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer srcFile.Close()
	//分割path目录
	destSplitPathDirs := strings.Split(dest, "/")
	//检测时候存在目录
	destSplitPath := ""
	for index, dir := range destSplitPathDirs {
		if index < len(destSplitPathDirs)-1 {
			destSplitPath = destSplitPath + dir + "/"
			b, _ := PathExists(destSplitPath)
			if b == false {
				//创建目录
				err := os.Mkdir(destSplitPath, os.ModePerm)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	dstFile, err := os.Create(dest)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	io.Copy(dstFile, srcFile)
	if fm != "" && to != "" {
		buf, _ := ioutil.ReadFile(dest)
		content := string(buf)
		//替换
		newContent := strings.Replace(content, fm, to, -1)

		//重新写入
		ioutil.WriteFile(dest, []byte(newContent), 0)
	}
	return
}
