package main

import (
	"fmt"
	"github.com/ota42y/resource-semaphore/resources"
)

func main() {
	dbs := []string{"host1", "host2", "host3"}

	sResources := resources.NewStringResources(dbs)

	// doesn't close channel
	r, _ := sResources.Withdraw()

	fmt.Println("resource ", r)

	sResources.Close() // fatal error: all goroutines are asleep - deadlock!
	fmt.Println("finish")
}
