package main

// 009、显示一个类型的方法集
func main() {
	// 我们的最后一个例子是使用 reflect.Type 来打印任意值的类型和枚举它的方法：
	// （见 files/methods/methods.go）

	// reflect.Type 和 reflect.Value 都提供了一个 Method 方法
	// 每次 t.Method(i) 调用将一个 reflect.Method 的实例，对应一个用于描述一个方法的名称和类型的结构体
	// 每次 v.Method(i) 方法调用都返回一个 reflect.Value 以表示对应的值（6.4 节），也就是一个方法是帮到它的接收者的
	// 使用 reflect.Value.Call 方法（在此就不列举了），将可以调用一个 Func 类型的 Value，但这个例子中只用到了它的类型

	// 这是属于 time.Duration 和 *strings.Replacer 两个类型的方法：
	// （见 files/methods/methods_test.go）
}
