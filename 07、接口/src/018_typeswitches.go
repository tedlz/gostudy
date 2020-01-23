package main

// 018、类型开关
func main() {
	// 接口被以两种不同的方式使用

	// 在第一个方式中，以 io.Reader、io.Writer、fmt.Stringer、sort.Interface、http.Handler 和 error 为典型
	// 一个接口的方法表达了实现这个接口的具体类型间的相似性，但是隐藏了代表的细节和这些具体类型本身的操作
	// 重点在于方法上，而不是本身的类型上

	// 第二个方式利用一个接口值可以持有各种具体类型值的能力并且将这个接口认为是这个类型的 union（联合）
	// 类型断言用来动态的区别这些类型并且对每一种情况都不一样
	// 在这个方式中，重点在于具体的类型满足这个接口，而不是在于接口的方法（如果它确实有一些的话），并且没有任何的信息隐藏
	// 我们将以这种方式使用的接口描述为 discriminated unions（可识别联合）

	// 如果你熟悉面向对象编程，
	// 你可能会将这两种方式当做是 subtype polymorphism（子类型多态）和 ad hoc polymorphism（非参数多态）
	// 但是你不需要去记住这些术语
	// 对于本章剩下的部分，我们将会呈现一些第二种方式的例子

	// 和其它那些语言一样，Go 语言查询一个 SQL 数据库的 API 会干净地将查询中固定的部分和变化的部分分开
	// 一个调用的例子可能看起来像这样：
	//
	// import "database/sql"
	// func listTracks(db sql.DB, artist string, minYear, maxYear int) {
	// 	   result, err := db.Exec(
	// 		   "SELECT * FROM tracks WHERE artist = ? AND ? <= year AND year <= ?",
	// 		   artist, minYear, maxYear)
	// 	   ......
	// }
	//
	// Exec 方法使用字面量替换在查询字符串中的每个 '?'
	// SQL 字面量表示相应参数的值，它有可能是一个布尔值、一个数字、一个字符串、或者 nil 空值
	// 用这种方式构造查询可以帮助避免 SQL 注入攻击，这种攻击就是对手可以通过利用输入内容中不正确的引文来控制查询语句
	// 在 Exec 函数内部，我们可能会找到像下面这样的一个函数，它会将每一个参数值转换成它的 SQL 字面量符号：
	//
	// func sqlQuote(x interface{}) string {
	// 	   if x == nil {
	// 		   return "NULL"
	// 	   } else if _, ok := x.(int); ok {
	// 		   return fmt.Sprintf("%d", x)
	// 	   } else if _, ok := x.(uint); ok {
	// 		   return fmt.Sprintf("%d", x)
	// 	   } else if b, ok := x.(bool); ok {
	// 		   if b {
	// 			   return "TRUE"
	// 		   }
	// 		   return "FALSE"
	// 	   } else if s, ok := x.(string); ok {
	// 		   return sqlQuoteString(x)
	// 	   } else {
	// 		   panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	// 	   }
	// }
	//
	// switch 语句可以简化 if-else 链，如果这个 if-else 链对一连串值做相等测试的话
	// 一个相似的 type switch（类型开关）可以简化类型断言的 if-else 链

	// 在它最简单的形式中，一个类型开关像普通的 switch 语句一样，它的运算对象是 x.(type)
	// - 它使用了关键词字面量 type，并且每个 case 有一到多个类型，一个类型开关基于这个接口值的动态类型使一个多路分支有效
	// 这个 nil 的 case 和 if x == nil 匹配，并且这个 default 的 case 和如果其它 case 都不匹配的情况匹配
	// 一个对 sqlQuote 的类型开关可能会有这些 case：
	//
	// switch x.(type) {
	// case nil:       // ...
	// case int, uint: // ...
	// case bool:      // ...
	// case string:    // ...
	// default:        // ...
	// }
	//
	// 和 §1.8 中普通的 switch 语句一样，每一个 case 会被顺序的进行考虑，并且当一个匹配找到时，这个 case 中的内容会被执行
	// 当一个或多个 case 类型是接口时，case 的顺序就会变得很重要，因为可能会有两个 case 同时匹配的情况
	// default case 相对其它 case 的位置是无所谓的，它不会允许落空发生

	// 注意到在原来的函数中，对于 bool 和 string 情况的逻辑需要通过类型断言访问提取的值
	// 因为这个做法很典型，类型开关语句有一个扩展的形式，它可以将提取的值绑定到一个在每个 case 范围内的新变量
	// switch x.(type) { // ... }

	// 这里我们已经将新的变量也命名为 x，和类型断言一样，重用变量名是很常见的
	// 和一个 switch 语句相似地，一个类型开关隐式的创建了一个语言块，因此新变量 x 的定义不会和外面块中的 x 变量冲突
	// 每一个 case 也会隐式的创建一个单独的语言块

	// 使用类型开关的扩展形式来重写 sqlQuote 函数会让这个函数更加的清晰：
	//
	// func sqlQuote(x interface{}) string {
	// 	   switch x := x.(type) {
	// 	   case nil:
	// 		   return "NULL"
	// 	   case int, uint:
	// 		   return fmt.Sprintf("%d", x)
	// 	   case bool:
	// 		   if x {
	// 			   return "TRUE"
	// 		   }
	// 		   return "FALSE"
	// 	   case string:
	// 		   return sqlQuoteString(x)
	// 	   default:
	// 		   panic("unexpected type %T: %v", x, x)
	// 	   }
	// }
	//
	// 在这个版本的函数中，在每一个单一类型的 case 内部，变量 x 和这个类型的 case 相同
	// 例如，变量 x 在 bool 的 case 中是 bool 类型，在 string 的 case 中是 string 类型
	// 在所有其它的情况中，变量 x 是 switch 运算对象的类型（接口），在这个例子中运算的对象是一个 interface{}
	// 当多个 case 需要相同的操作时，比如 int 和 uint 的情况，类型开关可以很容易的合并这些情况

	// 尽管 sqlQuote 接收一个任意类型的参数，但是这个函数只会在它的参数，匹配类型开关中的一个 case 时，运行到结束
	// 其它情况的它会 panic 出 unexpected type 消息
	// 虽然 x 的类型是 interface{}，但我们把它认为是一个 int/uint/bool/string/nil 值的 discriminated union 可识别联合
}
