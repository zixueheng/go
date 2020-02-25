# Go 二维码
`func WriteFile(content string, level RecoveryLevel, size int, filename string) error`

WriteFile 函数的原型定义如上，它有几个参数，大概意思如下：

- content 表示要生成二维码的内容，可以是任意字符串；
- level 表示二维码的容错级别，取值有 Low、Medium、High、Highest；
- size 表示生成图片的 width 和 height，像素单位；
- filename 表示生成的文件名路径；
- RecoveryLevel 类型其实是个 int，它的定义和常量如下：
```go
type RecoveryLevel int

const (
    // Level L: 7% error recovery.
    Low RecoveryLevel = iota

    // Level M: 15% error recovery. Good default choice.
    Medium

    // Level Q: 25% error recovery.
    High

    // Level H: 30% error recovery.
    Highest
)
```

## 生成二维码图片字节

有时候我们不想直接生成一个 PNG 文件存储，我们想对 PNG 图片做一些处理，比如缩放了，旋转了，或者网络传输了等，基于此，我们可以使用 Encode 函数，生成一个 PNG 图片的字节流，这样我们就可以进行各种处理了。

`func Encode(content string, level RecoveryLevel, size int) ([]byte, error)`

用法和 WriteFile 函数差不多，只不过返回的是一个 []byte 字节数组，这样我们就可以对这个字节数组进行处理了。

## 自定义二维码

除了以上两种快捷方式，go-qrcode 库还为我们提供了对二维码的自定义方式，比如我们可以自定义二维码的前景色和背景色等。qrcode.New 函数可以返回一个 *QRCode，我们可以对 *QRCode 设置，实现对二维码的自定义。

比如我们设置背景色为绿色，前景色为白色的二维码：

```go
    package main
    import(
        "github.com/skip2/go-qrcode"
        "image/color"
        "log"
    )
    func main() {
        qr,err:=qrcode.New("http://c.biancheng.net/",qrcode.Medium)
        if err != nil {
            log.Fatal(err)
        } else {
            qr.BackgroundColor = color.RGBA{50,205,50,255}
            qr.ForegroundColor = color.White
            qr.WriteFile(256,"./golang_qrcode.png")
        }
    }
```

