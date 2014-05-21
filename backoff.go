package backoff

import (
	"time"
)

type BackoffWaiter struct {
	a               int
	b               int
	maxWait         int
	waitCalledCount int
}

func NewBackoffWaiter(maxWait int) *BackoffWaiter {
	return &BackoffWaiter{0, 1, maxWait, 0}
}

func (bw *BackoffWaiter) Wait() {
	bw.waitCalledCount++

	bw.b, bw.a = bw.b+bw.a, bw.b

	wait := bw.b
	if wait > bw.maxWait {
		wait = bw.maxWait
	}

	<-time.After(time.Second * time.Duration(wait))
}

func (bw *BackoffWaiter) Reset() {
	bw.a = 0
	bw.b = 1
}

func (bw *BackoffWaiter) WaitCalledCount() int {
	return bw.waitCalledCount
}
