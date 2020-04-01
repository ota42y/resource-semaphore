package resources

import "github.com/ota42y/resource-semaphore/semaphore"

type StringResources struct {
	se *semaphore.ResourceSemaphore
	sArray []string
}

func NewStringResources(sArray []string) *StringResources {
	return &StringResources{
		se: semaphore.NewResourceSemaphore(len(sArray)),
		sArray: sArray,
	}
}

func (c *StringResources) Withdraw() (string, chan bool) {
	number, ch := c.se.Withdraw()

	return c.sArray[number], ch
}

func (c *StringResources) Close() {
	c.se.Close()
}
