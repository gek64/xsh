package main

import (
	"fmt"
	"gek_app"
	"os"
	"path/filepath"
)

type Resources struct {
	gek_app.Resources
}

// 安装资源(从本地文件解压,无需网络下载)
func (r Resources) installFromLocal(localFile string) (err error) {
	// 检查本地文件是否存在
	_, err = os.Stat(localFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s is not exist", localFile)
	}

	// 检测资源安装路径是否存在
	// 不存在则创建
	_, err = os.Stat(r.Location)
	if os.IsNotExist(err) {
		err = os.MkdirAll(r.Location, 0755)
		if err != nil {
			return err
		}
	}

	// 解压资源文件到资源安装路径
	err = gek_app.Extract(localFile, r.Files, r.Location)
	if err != nil {
		return err
	}

	// 赋权644
	for _, file := range r.Files {
		err = os.Chmod(filepath.Join(r.Location, file), 0644)
		if err != nil {
			return err
		}
	}

	return nil
}