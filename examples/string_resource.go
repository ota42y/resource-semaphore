package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ota42y/resource-semaphore/resources"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dbs := []string{"host1", "host2", "host3"}

	sResources := resources.NewStringResources(dbs)
	wg := &sync.WaitGroup{}
	for i := 1; i < 100; i++ {
		wg.Add(1)
		number := i
		go func() {
			time.Sleep(time.Duration(rand.Uint32() / 1000))
			workString(number, sResources)
			wg.Done()
		}()

	}
	wg.Wait()

	sResources.Close()
	fmt.Println("finish")
}

func workString(number int, c *resources.StringResources) {
	r, ch := c.Withdraw()
	fmt.Println("work ", number, " resource: ", r)
	close(ch)
}
