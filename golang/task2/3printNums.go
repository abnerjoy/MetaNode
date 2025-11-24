package main

import "fmt"

func PrintOddNums() {
	for i := 1; i < 10; i = i + 2 {
		fmt.Println(i)
	}
}
func PrintEvenNums() {
	for i := 2; i <= 10; i = i + 2 {
		fmt.Println(i)
	}
}

//func main() {
//	done := make(chan bool, 2)
//	go func() {
//		PrintOddNums()
//		done <- true
//	}()
//
//	go func() {
//		PrintEvenNums()
//		done <- true
//	}()
//	<-done
//	<-done
//}
