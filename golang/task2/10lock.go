package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
// 考察点 ：原子操作、并发数据安全。

func main() {
	var counter int64
	var wg sync.WaitGroup

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 每个协程进行1000次递增操作
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}

	// 等待所有协程完成
	wg.Wait()

	// 输出最终结果
	fmt.Println("计数器最终值:", atomic.LoadInt64(&counter))
}
