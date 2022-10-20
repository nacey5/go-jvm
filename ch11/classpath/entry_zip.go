package classpath

import "archive/zip"
import "errors"
import "io/ioutil"
import "path/filepath"

type ZipEntry struct {
	absPath string
	zipRC   *zip.ReadCloser
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absPath, nil}
}

func (this *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	if this.zipRC == nil {
		err := this.openJar()
		if err != nil {
			return nil, nil, err
		}
	}

	classFile := this.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}

	data, err := readClass(classFile)
	return data, this, err
}

// todo: close zip
func (this *ZipEntry) openJar() error {
	r, err := zip.OpenReader(this.absPath)
	if err == nil {
		this.zipRC = r
	}
	return err
}

func (this *ZipEntry) findClass(className string) *zip.File {
	for _, f := range this.zipRC.File {
		if f.Name == className {
			return f
		}
	}
	return nil
}

func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// read class data
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (this *ZipEntry) String() string {
	return this.absPath
}
