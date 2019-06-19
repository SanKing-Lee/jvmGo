package main

import (
	"fmt"
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
	// 输出假装运行了虚拟机
	fmt.Printf("classpath:%s class:%s args:%v\n", cmd.cpOption, cmd.class, cmd.args)
}
