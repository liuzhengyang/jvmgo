package classpath

import (
	"os"
	"path/filepath"
	"fmt"
	"errors"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClassPath(jreOption)
	cp.parseUserClassPath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClassPath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	fmt.Printf("jre lib path: %s\n", jreLibPath)
	self.bootClasspath = newWildcardEntry(jreLibPath)
	fmt.Println(jreLibPath)

	// jre/lib/ext/*
	jreExtpath := filepath.Join(jreDir, "lib", "ext", "*")
	fmt.Printf("jre ext lib path: %s\n", jreLibPath)
	self.extClasspath = newWildcardEntry(jreExtpath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		tmp := filepath.Join(jh, "jre")
		fmt.Printf("return javahome plus jre %s\n", tmp)
		return tmp
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true;
}

func (self *Classpath) parseUserClassPath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		if errors.New("error") == nil {
			fmt.Println("equals")
		}
		fmt.Printf("data %s\n", data)
		fmt.Printf("err %s\n", err)
		fmt.Printf("err %s\n", err == nil)
		b := (err == nil)
		fmt.Printf("b %s\n", b)
		fmt.Printf("b %s\n", errors.New("error"))
		fmt.Printf("b %s\n", (errors.New("error") == nil))
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}