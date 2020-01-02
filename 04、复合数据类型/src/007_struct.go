package main

import (
	"fmt"
	"time"
)

// Employee 定义一个结构
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

// Employee2 相同的类型可以被合并，例如以上可以写作
// 合并了同为 string 类型的 Name 和 Address
// 当然 Position 也是 string，同样可以合并，但那样就相当于定义了不同的结构体类型
// 因为在结构体中，结构体中成员的顺序也很重要
// 如果结构体中成员的名字以大写开头，那么该成员就是导出的，这是由 Go 语言导出规则决定的
// 一个结构体可能同时包含导出和未导出的成员
// 完整的结构体写法通常只在类型声明语句的地方出现
type Employee2 struct {
	ID            int
	Name, Address string
	DoB           time.Time
	Position      string
	Salary        int
	ManagerID     int
}

var dilbert Employee

// 007、结构体
// go run 007_struct.go
// 输出：
// -5000
// "Senior "
// Senior  (proactive team player)
// USA
// 0 100
func main() {
	// 结构体是一种聚合的数据类型，是由零个或多个任意类型的值聚合成的实体，每个值称为结构体的成员
	// dilbert 结构体变量的成员可以通过点操作符访问，例如 dilbert.Name 或 dilbert.DoB
	// 因为 dilbert 是一个变量，它所有的成员也同样是变量，我们可以直接对每个成员赋值
	dilbert.ID = 1
	dilbert.Salary -= 5000
	fmt.Println(dilbert.Salary) // -5000
	// 或是对成员取地址，然后通过指针访问
	position := &dilbert.Position
	*position = "Senior " + *position
	fmt.Printf("%q\n", dilbert.Position) // "Senior "
	// 点操作符也可以和指向结构体的指针一起工作
	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position += " (proactive team player)"
	// 相当于 (*employeeOfTheMonth).Position += " (proactive team player)"
	fmt.Println(dilbert.Position) // Senior  (proactive team player)

	// EmployeeByID 根据给定的员工 ID 返回对应的员工信息结构体的指针
	// 可以使用点操作符来访问里面的成员
	fmt.Println(EmployeeByID(a.ID).Position) // USA
	// 可以给里面的成员赋值
	EmployeeByID(b.ID).Salary = 100
	fmt.Println(EmployeeByID(a.ID).Salary, EmployeeByID(b.ID).Salary) // 0 100
}

var a = Employee{
	ID:       1,
	Name:     "a",
	Position: "USA",
}
var b = Employee{
	ID:       2,
	Name:     "b",
	Position: "UK",
}

// EmployeeByID 根据给定的员工 ID 返回对应的员工信息结构体的指针
func EmployeeByID(id int) *Employee {
	switch id {
	case a.ID:
		return &a
	case b.ID:
		return &b
	}
	return nil
}
