package main

func Pointer(p *int) {
	*p += 10
}

func Add(p int) {
	p += 10
}

//func main() {
//	a := 10
//	Pointer(&a)
//	fmt.Println(a)
//	b := 10
//	Add(10)
//	fmt.Println(b)
//}
