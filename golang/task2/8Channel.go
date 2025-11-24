package main

// 目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数
// ，消费者协程从通道中接收这些整数并打印。
// 考察点 ：通道的缓冲机制。

//func main() {
//	// 创建一个带缓冲的通道，缓冲区大小为10
//	ch := make(chan int, 10)
//	var wg sync.WaitGroup
//	wg.Add(2)
//
//	// 生产者协程
//	go func() {
//		defer wg.Done()
//		defer close(ch) // 生产完成后关闭通道
//		for i := 1; i <= 100; i++ {
//			ch <- i // 发送整数到通道
//		}
//	}()
//
//	// 消费者协程
//	go func() {
//		defer wg.Done()
//		for num := range ch { // 从通道接收整数
//			fmt.Println(num)
//		}
//	}()
//
//	wg.Wait() // 等待两个协程都完成
//}
