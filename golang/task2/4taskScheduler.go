package main

import (
	"sync"
	"time"
)

type Task func() error

func TaskScheduler(tasks []Task) []time.Duration {
	n := len(tasks)
	var durs = make([]time.Duration, n)
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i, t := range tasks {
		go func() {
			defer wg.Done()
			start := time.Now()
			_ = t()
			durs[i] = time.Since(start)
		}()
	}
	wg.Wait()
	return durs
}

//func main() {
//	tasks := []Task{
//		func() error {
//			time.Sleep(time.Millisecond * 500)
//			return nil
//		},
//		func() error {
//			time.Sleep(time.Millisecond * 200)
//			return nil
//		},
//		func() error {
//			time.Sleep(time.Millisecond * 300)
//			return nil
//		},
//	}
//	durs := TaskScheduler(tasks)
//	for _, d := range durs {
//		fmt.Println(d)
//	}
//}
