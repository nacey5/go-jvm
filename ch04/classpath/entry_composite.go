package classpath

import (
	"errors"
	"strings"
)

// CompositeEntry 合成Entry
type CompositeEntry []Entry

// CompositeEntry由更小的Entry组成，把路径分割之后构建多个实体Entry
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 每次都调用一个子方法
func (this CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range this {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (this CompositeEntry) String() string {
	strs := make([]string, len(this))
	for i, entry := range this {
		strs[i] = entry.String()
	}
	//将得到的字符串使用分隔符进行分隔
	return strings.Join(strs, pathListSeparator)
}
