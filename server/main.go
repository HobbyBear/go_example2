package main

import (
	"fmt"
	"go.uber.org/atomic"
	"time"
)

var (
	lock = tryLock{atomic.Int64{}}
)

type tryLock struct {
	atomic.Int64
}

func (t *tryLock) TryLock() bool {
	return t.CAS(0, 1)
}

func (t *tryLock) UnLock() {
	t.Store(0)
}

func main() {
	//res := lock.TryLock()
	//fmt.Println(res)
	//res = lock.TryLock()
	//fmt.Println(res)
	//res = lock.TryLock()
	//fmt.Println(res)
	//lock.UnLock()
	//fmt.Println(lock.Load())
	//res = lock.TryLock()
	//fmt.Println(res)
	fmt.Println(time.Now().Format("0102150405"))
	fmt.Println(time.Now().Unix())

}
