package main

import (
	"fmt"
	"go-jvm/ch03/classfile"
	"go-jvm/ch03/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	cp := classpath.Parse(cmd.xJreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printClassInfo(cf)
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classdata, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(classdata)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version:%v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count:%v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags:0x%x\n", cf.AccessFlags())
	fmt.Printf("this class:%v\n", cf.ClassName())
	fmt.Printf("super class:%v\n", cf.SuperClassName())
	fmt.Printf("interfaces:%v\n", cf.InterfaceNames())
	fmt.Printf("fields count:%v", len(cf.Fields()))
	for _, m := range cf.Fields() {
		fmt.Printf("\t%s\n", m.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
