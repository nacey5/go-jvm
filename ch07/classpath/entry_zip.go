package classpath

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("NewZipEntry error")
		panic(err)
	}
	return &ZipEntry{absPath}
}

// 从ZIP文件中提取class文件
func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//从路径读取一个zip文件
	r, err := zip.OpenReader(this.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			//读取文件中的所有字节
			data, err := io.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, this, nil
		}

	}
	return nil, nil, errors.New("class not found：" + className)
}

func (this *ZipEntry) String() string {
	return this.absPath
}
