package main

import (
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

type MyTest struct {
	param1 uint64
	_p1    [8]int64
	param2 uint64
	_p2    [8]int64
}

var addTimes = 100000000
var wg sync.WaitGroup

func Inc(num *uint64) {
	for i := 0; i < addTimes; i++ {
		atomic.AddUint64(num, 1)
	}
	wg.Done()
}

func BenchmarkTestProcessNum1(b *testing.B) {
	runtime.GOMAXPROCS(1)
	myTest := &MyTest{}
	wg.Add(2)
	go Inc(&myTest.param1)
	go Inc(&myTest.param2)
	wg.Wait()
}

func BenchmarkTestProcessNum2(b *testing.B) {
	runtime.GOMAXPROCS(2)
	myTest := &MyTest{}
	wg.Add(2)
	go Inc(&myTest.param1)
	go Inc(&myTest.param2)
	wg.Wait()
}
