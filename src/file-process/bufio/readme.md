# 使用buffer读取文件

buffer 是缓冲器的意思，Go语言要实现缓冲读取需要使用到 bufio 包。bufio 包本身包装了 io.Reader 和  io.Writer 对象，同时创建了另外的 Reader 和 Writer 对象，因此对于文本 I/O 来说，bufio 包提供了一定的便利性。

 buffer 缓冲器的实现原理就是，将文件读取进缓冲（内存）之中，再次读取的时候就可以避免文件系统的 I/O 从而提高速度。同理在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。

## 使用 bufio 包写入文件

bufio 和 io 包中有很多操作都是相似的，唯一不同的地方是 bufio 提供了一些缓冲的操作，如果对文件 I/O 操作比较频繁的，使用 bufio 包能够提高一定的性能。

 在 bufio 包中，有一个 Writer 结构体，而其相关的方法支持一些写入操作，如下所示。

```go
    // Writer 是一个空的结构体，一般需要使用 NewWriter 或者 NewWriterSize 来初始化一个结构体对象
    type Writer struct {
            // contains filtered or unexported fields
    }

	// NewWriterSize 和 NewWriter 函数:
    // 返回默认缓冲大小的 Writer 对象(默认是4096)
    func NewWriter(w io.Writer) *Writer
    // 指定缓冲大小创建一个 Writer 对象
    func NewWriterSize(w io.Writer, size int) *Writer

	// Writer 对象相关的写入数据的方法:
    // 把 p 中的内容写入 buffer，返回写入的字节数和错误信息。如果 nn < len(p)，返回错误信息中会包含为什么写入的数据比较短
    func (b *Writer) Write(p []byte) (nn int, err error)
    //将 buffer 中的数据写入 io.Writer
    func (b *Writer) Flush() error

	// 以下三个方法可以直接写入到文件中:
    // 写入单个字节
    func (b *Writer) WriteByte(c byte) error
    // 写入单个 Unicode 指针返回写入字节数错误信息
    func (b *Writer) WriteRune(r rune) (size int, err error)
    // 写入字符串并返回写入字节数和错误信息
    func (b *Writer) WriteString(s string) (int, error)
```

## 使用 bufio 包读取文件

使用 bufio 包读取文件也非常方便，我们先来看下 bufio 包的相关的 Reader 函数方法：

```go
    //首先定义了一个用来缓冲 io.Reader 对象的结构体，同时该结构体拥有以下相关的方法
    type Reader struct {
    }

    //NewReader 函数用来返回一个默认大小 buffer 的 Reader 对象（默认大小是 4096） 等同于 NewReaderSize(rd,4096)
    func NewReader(rd io.Reader) *Reader

    //该函数返回一个指定大小 buffer（size 最小为 16）的 Reader 对象，如果 io.Reader 参数已经是一个足够大的 Reader，它将返回该 Reader
    func NewReaderSize(rd io.Reader, size int) *Reader

    //该方法返回从当前 buffer 中能被读到的字节数
    func (b *Reader) Buffered() int

    //Discard 方法跳过后续的 n 个字节的数据，返回跳过的字节数。如果 0 <= n <= b.Buffered()，该方法将不会从 io.Reader 中成功读取数据
    func (b *Reader) Discard(n int) (discarded int, err error)

    //Peek 方法返回缓存的一个切片，该切片只包含缓存中的前 n 个字节的数据
    func (b *Reader) Peek(n int) ([]byte, error)

    //把 Reader 缓存对象中的数据读入到 []byte 类型的 p 中，并返回读取的字节数。读取成功，err 将返回空值
    func (b *Reader) Read(p []byte) (n int, err error)

    //返回单个字节，如果没有数据返回 err
    func (b *Reader) ReadByte() (byte, error)

    //该方法在 b 中读取 delimz 之前的所有数据，返回的切片是已读出的数据的引用，切片中的数据在下一次的读取操作之前是有效的。如果未找到 delim，将返回查找结果并返回 nil 空值。因为缓存的数据可能被下一次的读写操作修改，因此一般使用 ReadBytes 或者 ReadString，他们返回的都是数据拷贝
    func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

    //功能同 ReadSlice，返回数据的拷贝
    func (b *Reader) ReadBytes(delim byte) ([]byte, error)

    //功能同 ReadBytes，返回字符串
    func (b *Reader) ReadString(delim byte) (string, error)

    //该方法是一个低水平的读取方式，一般建议使用 ReadBytes('\n') 或 ReadString('\n')，或者使用一个 Scanner 来代替。ReadLine 通过调用 ReadSlice 方法实现，返回的也是缓存的切片，用于读取一行数据，不包括行尾标记（\n 或 \r\n）
    func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

    //读取单个 UTF-8 字符并返回一个 rune 和字节大小
    func (b *Reader) ReadRune() (r rune, size int, err error)
```

