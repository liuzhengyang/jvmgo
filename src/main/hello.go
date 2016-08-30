package main

import (
	"fmt"
	"os"
	"path/filepath"
	"errors"
)

func main() {
	fmt.Println("Hello, World")
	javaHome := os.Getenv("JAVA_HOME")
	path := filepath.Join(javaHome, "jre", "lib")
	fmt.Println(path)
	err := errors.New("nihao")
	fmt.Println(err == nil)
}