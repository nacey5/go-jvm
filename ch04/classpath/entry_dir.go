package classpath

import (
	"fmt"
	"os"
	"path/filepath"
)

// DirEntry 用于存放目录绝对路径
type DirEntry struct {
	absDir string
}

// 工厂实例方法
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		fmt.Println("newDirEntry error")
		panic(err)
	}
	return &DirEntry{absDir: absDir}
}

// 把目录名和class文件名拼接成一个完整的路径,再读取class中的内容
func (this *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(this.absDir, className)
	data, err := os.ReadFile(fileName)
	return data, this, err
}

func (this *DirEntry) String() string {
	return this.absDir
}
