package main

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkVal(b *testing.B) {
	ch, wg := goPool(goPoolSize)
	for i := 0; i < b.N; i++ {
		ch <- func() {
			var a TVal
			a.i++
			ptr = &a.i // to make a escape
		}
	}
	close(ch)
	wg.Wait()
}

func BenchmarkGoLocal(b *testing.B) {
	ch, wg := goPool(goPoolSize)
	for i := 0; i < b.N; i++ {
		ch <- func() {
			go_local a TVal
			a.i++
			ptr = &a.i
		}
	}
	close(ch)
	wg.Wait()
}

func BenchmarkGoLocal2(b *testing.B) {
	ch, wg := goPool(goPoolSize)
	for i := 0; i < b.N; i++ {
		ch <- func() {
			a, _ := runtime.NewGoLocal[TVal](1, func() TVal {
				return TVal{}
			})
			a.Val.i++
			ptr = &a.Val.i
		}
	}
	close(ch)
	wg.Wait()
}

func BenchmarkPoolVal(b *testing.B) {
	ch, wg := goPool(goPoolSize)
	for i := 0; i < b.N; i++ {
		ch <- func() {
			a := pool.Get().(*TVal)
			a.i++
			ptr = &a.i
			pool.Put(a)
		}
	}
	close(ch)
	wg.Wait()
}

func goPool(workers int) (chan func(), *sync.WaitGroup) {
	wg := &sync.WaitGroup{}
	ch := make(chan func(), workers)
	for j := 0; j < workers; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range ch {
				f()
			}
		}()
	}
	return ch, wg
}

var ptr *int

type TVal struct {
	bs [1024]byte
	i  int
}

var pool = sync.Pool{New: func() interface{} {
	return &TVal{}
},
}

var goPoolSize = 100
