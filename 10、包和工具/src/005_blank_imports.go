package main

// 005、包的匿名导入
func main() {

}

// 如果只是导入一个包而并不使用导入的包将会导致一个编译错误
// 但是有时候我们只是想利用导入包而产生的副作用：
// 它会计算包级变量的初始化表达式和执行导入包的 init 初始化函数（见 2.6.2 节）
// 这时候我们需要抑制 unused import 编译错误，我们可以用下划线 _ 来重命名导入的包
// 像往常一样，下划线 _ 为空白标识符，并不能被访问

// import _ "image/png" // 注册 png 解码器

// 这种方式叫做包的匿名导入
// 它通常是用来实现一个编译时机制，然后通过在 main 主程序入口选择性地导入附加的包
// 首先，让我们看看如何使用该特性，然后再看看它是如何工作的

// 标准库的 image 图像包包含了一个 Decode 函数，用于从 io.Reader 接口读取数据并解码图像，
// 它调用底层注册的图像解码器来完成任务，然后返回 image.Image 类型的图像
// 使用 image.Decode 很容易编写一个图像格式的转换工具，读取一种格式的图像，然后编码为另一种图像格式：
// （见 files/jpeg）

// 如果我们将 gopl.io/ch3/mandelbrot（3.3 节）的输出导入到这个程序的标准输入
// 它将解码输入的 png 格式图像，然后转换为 jpeg 格式的图像输出（图 3.3）

// $ go build gopl.io/ch3/mandelbrot
// $ go build gopl.io/ch10/jpeg
// $ ./mandelbrot | ./jpeg >mandelbrot.jpg
// Input format = png

// 要注意 image/png 包的匿名导入语句
// 如果没有这一行语句，程序依然可以编译和运行，但是它将不能正确识别和解码 png 格式的图像

// $ go build gopl.io/ch3/mandelbrot
// $ go build gopl.io/ch10/jpeg
// $ ./mandelbrot | ./jpeg >mandelbrot.jpg
// jpeg: image: unknown format

// 下面的代码演示了它的工作机制
// 标准库还提供了 gif、png 和 jpeg 等格式图像的解码器
// 用户也可以提供自己的解码器，但是为了保持程序体积较小，很多解码器并没有被全部包含，除非是明确需要支持的格式
// image.Decode 函数在解码时会依次查询支持的格式列表

// 每个格式驱动列表的每个入口指定了四件事情：
//   - 格式的名称
//   - 一个用于描述这种图像数据开头部分模式的字符串，用于解码器检测识别
//   - 一个 Decode 函数用于完成解码图像工作
//   - 一个 DecodeConfig 函数用于解码图像的大小和颜色空间的信息

// 每个驱动入口是通过 image.RegisterFormat 函数注册，一般是在每个格式包的 init 初始化函数中调用，
// 例如 image 包是这样注册的：

// package png
//
// func Decode(r io.Reader) (image.Image, error)
// func DecodeConfig(r io.Reader) (image.Config, error)
//
// func init() {
// 	   const pngHeader = "\x89PNG\r\n\x1a\n"
// 	   image.RegisterFormat("png", pngHeader, Decode, DecodeConfig)
// }

// 最终的效果是，主程序只需要匿名导入特定图像驱动包就可以用 image.Decode 解码对应格式的图像了
// 数据库包 database/sql 也是采用了类似的技术，让用户可以根据自己需要选择导入必要的数据库驱动
// 例如：

// import (
// 	   "database/sql"
// 	   _ "github.com/lib/pq" // 启动对 PostgreSQL 的支持
// 	   _ "github.com/go-sql-driver/mysql" // 启动对 MySQL 的支持
// )

// db, err = sql.Open("postgres", dbname) // OK
// db, err = sql.Open("mysql", dbname)    // OK
// db, err = sql.Open("sqlite3", dbname)  // returns error: unknown driver "sqlite3"
