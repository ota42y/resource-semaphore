package semaphore

type ResourceSemaphore struct {
	size    int
	storage chan int

	received chan int
	stopChan chan bool
	endChan  chan bool
}

func NewResourceSemaphore(size int) *ResourceSemaphore {
	a := &ResourceSemaphore{
		size:    size,
		storage: make(chan int, size),
		received: make(chan int, size),

		stopChan: make(chan bool, 1),
		endChan:  make(chan bool, 1),
	}

	for i := 0; i < size; i++ {
		a.storage <- i // init resources
	}

	go a.start()

	return a
}

func (c *ResourceSemaphore) start() {
	for {
		select {
		case r := <-c.received: // release resource so we add again
			c.storage <- r
		case <-c.stopChan:
			c.stop()
			return
		}
	}
}

// wait all resource released
func (c *ResourceSemaphore) stop() {
	count := 0

	for {
		select {
		case d := <-c.received: // wait for all resource released
			c.storage <- d
		case <-c.storage:
			count += 1
			if count == c.size {
				close(c.endChan)  // all resource released
				return
			}
		}
	}
}

// Withdraw return resource number and channel for release
// Please close channel when the withdrew resource released
func (c *ResourceSemaphore) Withdraw() (int, chan bool) {
	releasedChan := make(chan bool, 1)
	resource := <-c.storage

	go func() {
		<-releasedChan
		c.received <- resource
	}()

	return resource, releasedChan
}

// Close remove all resources
// don't Withdraw after close
func (c *ResourceSemaphore) Close() {
	close(c.stopChan)
	<-c.endChan
}


