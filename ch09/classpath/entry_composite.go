package classpath

import "errors"
import "strings"

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}

	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}

	return compositeEntry
}

func (this CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range this {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}

	return nil, nil, errors.New("class not found: " + className)
}

func (this CompositeEntry) String() string {
	strs := make([]string, len(this))

	for i, entry := range this {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}
