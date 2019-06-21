package main

import (
	"fmt"
	"strings"

	"github.com/sanking-lee/jvmGo/ch02/classpath"
)

func main() {
	cmd := parseCmd()    // 获取命令
	if cmd.versionFlag { // 显示版本信息
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" { // 打印提示信息
		printUsage()
	} else { // 运行虚拟机
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	// 输出假装运行了虚拟机
	fmt.Printf("classpath:%v class:%s args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
