# Go语言go mod包依赖管理工具
最早的时候，Go语言所依赖的所有的第三方库都放在 GOPATH 这个目录下面，这就导致了同一个库只能保存一个版本的代码。如果不同的项目依赖同一个第三方的库的不同版本，应该怎么解决？  

go module 是Go语言从 1.11 版本之后官方推出的版本管理工具，并且从 Go1.13 版本开始，go module 成为了Go语言默认的依赖管理工具。  

## 如何使用 Modules？
1) 首先需要把 golang 升级到 1.11 版本以上（现在 1.13 已经发布了，建议使用 1.13）。

2) 设置 GO111MODULE。

#### GO111MODULE
在Go语言 1.12 版本之前，要启用 go module 工具首先要设置环境变量 GO111MODULE，不过在Go语言 1.13 及以后的版本则不再需要设置环境变量。通过 GO111MODULE 可以开启或关闭 go module 工具。

- GO111MODULE=off 禁用 go module，编译时会从 GOPATH 和 vendor 文件夹中查找包；
- GO111MODULE=on 启用 go module，编译时会忽略 GOPATH 和 vendor 文件夹，只根据 go.mod下载依赖；
- GO111MODULE=auto（默认值），当项目在 GOPATH/src 目录之外，并且项目根目录有 go.mod 文件时，开启 go module。

Windows 下开启 GO111MODULE 的命令为：
set GO111MODULE=on 或者 set GO111MODULE=auto

MacOS 或者 [Linux](http://c.biancheng.net/linux_tutorial/) 下开启 GO111MODULE 的命令为：
export GO111MODULE=on 或者 export GO111MODULE=auto

在开启 GO111MODULE 之后就可以使用 go module 工具了，也就是说在以后的开发中就没有必要在 GOPATH 中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

常用的`go mod`命令如下表所示：

| 命令            | 作用                                           |
| --------------- | ---------------------------------------------- |
| go mod download | 下载依赖包到本地（默认为 GOPATH/pkg/mod 目录） |
| go mod edit     | 编辑 go.mod 文件                               |
| go mod graph    | 打印模块依赖图                                 |
| go mod init     | 初始化当前文件夹，创建 go.mod 文件             |
| go mod tidy     | 增加缺少的包，删除无用的包                     |
| go mod vendor   | 将依赖复制到 vendor 目录下                     |
| go mod verify   | 校验依赖                                       |
| go mod why      | 解释为什么需要依赖                             |

#### GOPROXY
Windows 下设置 GOPROXY 的命令为：
go env -w GOPROXY=https://goproxy.cn,direct

MacOS 或 Linux 下设置 GOPROXY 的命令为：
export GOPROXY=https://goproxy.cn

设置后可执行 go env 查看是否成功

#### 使用go get命令下载指定版本的依赖包
执行`go get `命令，在下载依赖包的同时还可以指定依赖包的版本。

- 运行`go get -u`命令会将项目中的包升级到最新的次要版本或者修订版本；
- 运行`go get -u=patch`命令会将项目中的包升级到最新的修订版本；
- 运行`go get [包名]@[版本号]`命令会下载对应包的指定版本或者将对应包升级到指定的版本。

提示：`go get [包名]@[版本号]`命令中版本号可以是 x.y.z 的形式，例如 go get foo@v1.2.3，也可以是 git 上的分支或 tag，例如 go get foo@master，还可以是 git 提交时的哈希值，例如 go get foo@e3702bed2。

## 如何在项目中使用

#### 创建一个新项目：

1) 在 GOPATH 目录之外新建一个目录，并使用`go mod init`初始化生成 go.mod 文件:
go mod init hello
go: creating new go.mod: module hello

go.mod 文件一旦创建后，它的内容将会被 go toolchain 全面掌控，go toolchain 会在各类命令执行时，比如`go get`、`go build`、`go mod`等修改和维护 go.mod 文件。

go.mod 提供了 module、require、replace 和 exclude 四个命令：

- module 语句指定包的名字（路径）；
- require 语句指定的依赖项模块；
- replace 语句可以替换依赖项模块；
- exclude 语句可以忽略依赖项模块。

初始化生成的 go.mod 文件如下所示：

module hello
go 1.13

2) 添加依赖。
新建一个 main.go 文件，写入以下代码：
```go
    package main
    import (
        "net/http"
        "github.com/labstack/echo"
    )
    func main() {
        e := echo.New()
        e.GET("/", func(c echo.Context) error {
            return c.String(http.StatusOK, "Hello, World!")
        })
        e.Logger.Fatal(e.Start(":1323"))
    }
```
执行`go run main.go`运行代码会发现 go mod 会自动查找依赖自动下载

go 会自动生成一个 go.sum 文件来记录 dependency tree

再次执行脚本`go run main.go`发现跳过了检查并安装依赖的步骤。

可以使用命令`go list -m -u all`来检查可以升级的 package，使用`go get -u need-upgrade-package`升级后会将新的依赖版本更新到 go.mod * 也可以使用`go get -u`升级所有依赖。

#### 使用内部包

增加文件 api/api.go

```go
    package api
    import (
        "net/http"
        "github.com/labstack/echo"
    )
    func HelloWorld(c echo.Context) error {
        return c.JSON(http.StatusOK, "hello world")
    }
```



运行`go run main.go` 会出现 `build _/D_/code/src/api: cannot find module for path _/D_/code/src/api`的错误

这是因为 main.go 中使用 internal package 的方法跟以前已经不同了，由于 go.mod 会扫描同工作目录下所有  package 并且变更引入方法，必须将 hello 当成路径的前缀，也就是需要写成 import hello/api（包的 模块内的绝对路径来导入，就是要从go.mod所在的目录开始），以往  GOPATH/dep 模式允许的 import ./api 已经失效。

main.go 导入api 的方式 改成 `import api "hello/api"` 再次运行`go run main.go` 就正常了。

下面就跟正常进行其他开发了。

## 使用 replace 替换无法直接获取的 package

由于某些已知的原因，并不是所有的 package 都能成功下载，比如：golang.org 下的包。

 modules 可以通过在 go.mod 文件中使用 replace 指令替换成 github 上对应的库，比如：

replace (
   golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a
 )

或者

replace golang.org/x/crypto v0.0.0-20190313024323-a1f597ede03a => github.com/golang/crypto v0.0.0-20190313024323-a1f597ede03a