package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Counter struct {
	result int
	sync.RWMutex
	ch chan int
}

// 线程饥饿
func (c *Counter) add(a int, wg *sync.WaitGroup) {
	c.Lock()
	defer c.Unlock()
	c.result += a
	wg.Done()
}

func (c Counter) get() int {
	c.Lock()
	defer c.Unlock()
	return c.result
}

func TestOpt(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go c.add(1, &wg)

	}
	wg.Wait()
	fmt.Println(c.get())
}
// 管道
func (c *Counter) add1(a int) {
	c.ch <- a
}
func (c *Counter) get1() int {
	return c.result
}
func (c *Counter) volatile1() {
	func() {
		for {
			c.result += <-c.ch
		}
	}()
}

func TestOpt1(t *testing.T) {
	c := &Counter{
		ch: make(chan int),
	}
	for i := 0; i < 100; i++ {
		go c.add1(1)
	}
	go c.volatile1()
	time.Sleep(1*time.Second)
	fmt.Println(c.get1())
}
// atomic
func (c *Counter) add2(a int , wg *sync.WaitGroup) {
	intSize := 32 << (^uint(0) >> 63) // 32 or 64
	if intSize == 64 {
		i := int64(c.result)
		c.result = int(atomic.AddInt64(&i, int64(a)))
	} else if intSize == 32 {
		i := int32(c.result)
		c.result = int(atomic.AddInt32(&i, int32(a)))
	}
	wg.Done()
}

func (c Counter) get2() int {
	return c.result
}

func TestOpt2(t *testing.T) {
	c := &Counter{}
	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go c.add2(1, &wg)
	}
	wg.Wait()
	fmt.Println(c.get2())
}