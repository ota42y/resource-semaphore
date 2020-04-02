package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ota42y/resource-semaphore/semaphore"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := semaphore.NewResourceSemaphore(10)
	wg := &sync.WaitGroup{}
	for i := 1; i < 100; i++ {
		wg.Add(1)
		number := i
		go func() {
			time.Sleep(time.Duration(rand.Uint32() / 1000))
			work(number, c)
			wg.Done()
		}()

	}
	wg.Wait()

	c.Close()
	fmt.Println("finish")
}

func work(number int, c *semaphore.ResourceSemaphore) {
	r, ch := c.Withdraw()
	fmt.Println("work ", number, " resource: ", r)
	close(ch)
}
