/*
	对命令行进行解析
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

/*
 *@description: 命令行结构体
 *@version:
 0.0.2 添加xjreoption属性
 0.0.1 初始化
*/
type Cmd struct {
	helpFlag    bool     // 是否为帮助命令行
	versionFlag bool     // 是否为版本命令行
	cpOption    string   // classpath可选项
	XjreOption  string   // 指定jre目录
	class       string   // 类名
	args        []string // 参数
}

/*
 *@description: 利用flag解析命令行，设置Cmd结构体系的各项属性
 *@params:
 *@return: 解析完成后的Cmd指针
 *@version:
 0.0.2  添加对XjreOption的解析
 0.0.1	初始化
*/
func parseCmd() *Cmd {
	// 实例化命令
	cmd := &Cmd{}
	// 把命令转化为各项属性
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

/*
 *@description: 打印用法说明
 *@params:
 *@return:
 *@version:
  0.0.1 初始化
*/
func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
