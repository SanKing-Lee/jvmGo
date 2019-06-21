package main

import (
	"flag"
	"fmt"
	"os"
)

// 命令行结构体
type Cmd struct {
	helpFlag    bool	// 是否为帮助命令行
	versionFlag bool	// 是否为版本命令行
	cpOption    string	// classpath可选项
	class       string	// 类名
	args        []string	// 参数
}

func parseCmd() *Cmd {
	// 实例化命令
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
