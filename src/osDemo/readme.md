# Golang os包说明

Go语言的 os 包中提供了操作系统函数的接口，是一个比较重要的包。顾名思义，os 包的作用主要是在服务器上进行系统的基本操作，如文件操作、目录操作、执行命令、信号与中断、进程、系统状态等等。  

#### 1) Hostname
函数定义:
`go func Hostname() (name string, err error) `
Hostname 函数会返回内核提供的主机名。

#### 2) Environ
函数定义:
` func Environ() []string `
Environ 函数会返回所有的环境变量，返回值格式为“key=value”的字符串的切片拷贝。

#### 3) Getenv
函数定义:
`func Getenv(key string) string`
Getenv 函数会检索并返回名为 key 的环境变量的值。如果不存在该环境变量则会返回空字符串。

#### 4) Setenv
函数定义:
` func Setenv(key, value string) error go`
Setenv 函数可以设置名为 key 的环境变量，如果出错会返回该错误。

#### 5) Exit
函数定义:
`func Exit(code int)`
Exit 函数可以让当前程序以给出的状态码 code 退出。一般来说，状态码 0 表示成功，非 0 表示出错。程序会立刻终止，并且 defer 的函数不会被执行。

#### 6) Getuid
函数定义:
`func Getuid() int`
Getuid 函数可以返回调用者的用户 ID。

#### 7) Getgid
函数定义:
`func Getgid() int`
Getgid 函数可以返回调用者的组 ID。

#### 8) Getpid
函数定义:
`func Getpid() int`
Getpid 函数可以返回调用者所在进程的进程 ID。

#### 9) Getwd
函数定义:
`func Getwd() (dir string, err error)`
Getwd 函数可以返回一个对应当前工作目录的根路径。如果当前目录可以经过多条路径抵达（因为硬链接），Getwd 会返回其中一个。

#### 10) Mkdir
函数定义:
`func Mkdir(name string, perm FileMode) error`
Mkdir 函数可以使用指定的权限和名称创建一个目录。如果出错，会返回 *PathError 底层类型的错误。

#### 11) MkdirAll
函数定义:
`func MkdirAll(path string, perm FileMode) error`
MkdirAll 函数可以使用指定的权限和名称创建一个目录，包括任何必要的上级目录，并返回 nil，否则返回错误。权限位 perm 会应用在每一个被该函数创建的目录上。如果 path 指定了一个已经存在的目录，MkdirAll 不做任何操作并返回 nil。

#### 12) Remove
函数定义:
`func Remove(name string) error`
Remove 函数会删除 name 指定的文件或目录。如果出错，会返回 *PathError 底层类型的错误。
RemoveAll 函数跟 Remove 用法一样，区别是会递归的删除所有子目录和文件。