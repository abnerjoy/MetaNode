package main

func SlicePointer(p *[]int) {
	for i := 0; i < len(*p); i++ {
		(*p)[i] = (*p)[i] * 2
	}
}

//func main() {
//	slice := []int{1, 2, 3, 4, 5}
//	SlicePointer(&slice)
//	fmt.Println(slice)
//}
