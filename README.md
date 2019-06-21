##用GO语言写JAVA虚拟机
参考书籍：《自己动手写Java虚拟机》
作者：张秀宏
出版社：机械工业出版社

**2019.06.22**
- ch02: 加载classpath
    - classpath
        - Classpath.go  类路径主要分为启动类，扩展类和用户类三种
        - entry.go  定义了所有用于加载类的入口的接口
            - readClass
            - String
        - entry_zip.go  对jar文件进行解析
        - entry_composite.go 对多个组合路径进行解析
        - entry_dir.go  对目录路径进行解析
        - entry_wildcard.go 对通配符进行解析
    - cmd.go    添加Xjreoption，用于解析传入命令中的-Xjre及其对应的jre路径
    - main.go   修改startJVM，测试classpath

**2019.06.19**
- ch01: 命令行工具
    - cmd.go
        - Cmd结构体
        - parseCmd()
        - printUsage()
    - main.go
        - main
        - startJVM(cmd *Cmd)
        
