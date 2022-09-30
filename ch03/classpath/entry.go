package classpath

import (
	"os"
	"strings"
)

// 定义分隔符号为 ;
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".ZIP") || strings.HasSuffix(path, "zip") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
