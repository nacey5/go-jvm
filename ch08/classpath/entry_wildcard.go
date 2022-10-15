package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	//去掉路径末尾的*号
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		//如果出现错误，立马返回
		if err != nil {
			return err
		}
		//如果传入的文件是一个文件且不是完整路径,跳过此次扫描
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		//如果此entry是以jar结尾的文件，加入compositeEntry
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}

	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
