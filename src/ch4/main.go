package main

import (
	"fmt"
)
import (
	"ch2/classpath"
	"ch3/classfile"
	"ch4/rtda"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	frame := rtda.NewFrame(100, 100)
	testLocalValr(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}
func testOperandStack(stack *rtda.OperandStack) {
	stack.PushInt(100)
	stack.PushInt(-100)
	stack.PushLong(102425325235)
	stack.PushLong(-10002030230520350)
	stack.PushFloat(3.2342342)
	stack.PushDouble(-23.23234235)
	stack.PushRef(nil)

	println(stack.PopRef())
	println(stack.PopDouble())
	println(stack.PopFloat())
	println(stack.PopLong())
	println(stack.PopLong())
	println(stack.PopInt())
	println(stack.PopInt())
}
func testLocalValr(vars rtda.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 1002121242142)
	vars.SetLong(4, -20232352323)
	vars.SetFloat(6, 2.3234)
	vars.SetDouble(7, 2.23823423223)
	vars.SetRef(9, nil)
	println(vars.GetInt(0))
	println(vars.GetInt(1))
	println(vars.GetLong(2))
	println(vars.GetLong(4))
	println(vars.GetFloat(6))
	println(vars.GetDouble(7))
	println(vars.GetRef(9))
	println("line")
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constants count: %v\n", len(cf.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", cf.AccessFlags())
	fmt.Printf("this class: %v\n", cf.ClassName())
	fmt.Printf("super class: %v\n", cf.SuperClassName())
	fmt.Printf("interfaces: %v\n", cf.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("\t0x%x\t%s\t%s\n", f.AccessFlags(), f.Descriptor(), f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("\t0x%x\t%s\t%s\n", m.AccessFlags(), m.Descriptor(), m.Name())
	}
}
