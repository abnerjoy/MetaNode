package main

/*题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。
在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。*/

type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
}

func (r Rectangle) Area() {
	println("Area Rectangle")
}

func (r Rectangle) Perimeter() {
	println("Perimeter Rectangle")
}

type Circle struct {
}

func (c Circle) Area() {
	println("Area Circle")
}

func (c Circle) Perimeter() {
	println("Perimeter Circle")
}

//
//func main() {
//	var rshape Shape = Rectangle{}
//	rshape.Area()
//	rshape.Perimeter()
//	var cshape Shape = Circle{}
//	cshape.Area()
//	cshape.Perimeter()
//
//}
