package classpath

import (
	"os"
	"path/filepath"
)

// Classpath 类加载路径
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

// 使用-Xjre选项解析启动类路径和扩展类路径
// 使用-classpath/-cp 选项解析用户类路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

// 如果用户没有提供-classpath/-cp，则使用当前目录作为用户类路径
func (this *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := this.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := this.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return this.userClasspath.readClass(className)
}

func (this *Classpath) String() string {
	return this.userClasspath.String()
}

// 解析启动类和扩展类路径
func (this *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	//终止 jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	this.bootClasspath = newWildcardEntry(jreLibPath)

	//终止 jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	this.extClasspath = newWildcardEntry(jreExtPath)
}

// 解析用户类路径
func (this *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	this.userClasspath = newEntry(cpOption)
}

// 优先用户输入的XJre选项进行jre目录
// 没有该选项,则在当前目录寻找
// 如果找不到，使用JAVA_HOME环境变量寻找
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	//使用os.Getenv()检索系统环境变量名
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

// 判读目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
