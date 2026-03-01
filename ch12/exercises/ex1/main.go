package main

import "sync"

func main() {
	ch := make(chan int, 20)
	var wg sync.WaitGroup
	wg.Add(2)
	for i := 0; i < 2; i++ {
		go func() {
			defer wg.Done()
			for i := range 10 {
				ch <- i
			}
		}()
	}

	var readerWg sync.WaitGroup
	readerWg.Add(1)
	go func() {
		defer readerWg.Done()
		for v := range ch {
			println(v)
		}
	}()

	wg.Wait()
	close(ch)
	readerWg.Wait()
}
