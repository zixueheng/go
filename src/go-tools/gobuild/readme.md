## go build 无参数编译

如果源码中没有依赖 GOPATH 的包引用，那么这些源码可以使用无参数 go build

`go build`

## go build+文件列表

编译同目录的多个源码文件时，可以在 go build 的后面提供多个文件名，go build 会编译这些源码，输出可执行文件，“go build+文件列表”的格式如下：

`go build file1.go file2.go……`

用“go build+文件列表”方式编译时，可执行文件默认选择文件列表中第一个源码文件作为可执行文件名输出。

如果需要指定输出可执行文件名，可以使用`-o`参数，参见下面的例子：

`go build -o myexec main.go lib.go`

上面代码中，在 go build 和文件列表之间插入了`-o myexec`参数，表示指定输出文件名为 myexec

使用“go build+文件列表”编译方式编译时，文件列表中的每个文件必须是同一个包的 Go 源码。也就是说，不能像 C++ 一样将所有工程的 Go 源码使用文件列表方式进行编译。编译复杂工程时需要用“指定包编译”的方式。

 “go build+文件列表”方式更适合使用Go语言编写的只有少量文件的工具。

## go build+包

“go build+包”在设置 GOPATH 后，可以直接根据包名进行编译，即便包内文件被增（加）删（除）也不影响编译指令。

## go build 编译时的附加参数

| 附加参数 | 备  注                                      |
| -------- | ------------------------------------------- |
| -v       | 编译时显示包名                              |
| -p n     | 开启并发编译，默认情况下该值为 CPU 逻辑核数 |
| -a       | 强制重新构建                                |
| -n       | 打印编译时会用到的所有命令，但不真正执行    |
| -x       | 打印编译时会用到的所有命令                  |
| -race    | 开启竞态检测                                |

## Golang交叉编译平台的二进制文件
```shell
# mac上编译linux和windows二进制
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build 
 
# linux上编译mac和windows二进制
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build 
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
 
# windows上编译mac和linux二进制
SET CGO_ENABLED=0 SET GOOS=darwin SET GOARCH=amd64 go build main.go
SET CGO_ENABLED=0 SET GOOS=linux SET GOARCH=amd64 go build main.go
```

## linux 上运行二进制文件
```shell
# 修改权限命令
chmod 777 程序名称

# 后台运行的命令
nohup ./程序名 & 

# 不输出错误信息
nohup ./程序名 >/dev/null 2>&1 &

# 如果要关闭程序，可以使用命令”ps” 查看后台程序的pid，然后使用“kill 程序pid”命令，关闭程序比如程序名为test，可以用如下命令查询
ps aux|grep test
```

#### 命令 nohup 和 & 
&：指在后台运行
nohup：不挂断的运行，注意并没有后台运行的功能，，就是指，用nohup运行命令可以使命令永久的执行下去，和用户终端没有关系，例如我们断开SSH连接都不会影响他的运行，注意了nohup没有后台运行的意思；&才是后台运行

&：是指在后台运行，但当用户退出(挂起)的时候，命令自动也跟着退出
那么，我们可以巧妙的吧他们结合起来用就是
nohup COMMAND &
这样就能使命令永久的在后台执行
例如：
1. sh test.sh &  
将sh test.sh任务放到后台 ，即使关闭xshell退出当前session依然继续运行，但标准输出和标准错误信息会丢失（缺少的日志的输出）

将sh test.sh任务放到后台 ，关闭xshell，对应的任务也跟着停止。
2. nohup sh test.sh  
将sh test.sh任务放到后台，关闭标准输入，终端不再能够接收任何输入（标准输入），重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行。
3. nohup sh test.sh  & 
将sh test.sh任务放到后台，但是依然可以使用标准输入，终端能够接收任何输入，重定向标准输出和标准错误到当前目录下的nohup.out文件，即使关闭xshell退出当前session依然继续运行