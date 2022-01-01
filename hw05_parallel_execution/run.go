package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	flag := true
	var fails int32
	chTask := make(chan Task)
	chBuff := make(chan int, n)
	for i := 0; i < n; i++ {
		chBuff <- 1
	}
	go func(fg *bool) {
		for _, task := range tasks {
			<-chBuff
			if *fg {
				chTask <- task
			} else {
				close(chTask)
				return
			}
		}
		close(chTask)
	}(&flag)
	for task := range chTask {
		task := task
		if fails < int32(m) {
			wg.Add(1)
			go func(wg *sync.WaitGroup) {
				defer wg.Done()
				err := task()
				if err != nil {
					atomic.AddInt32(&fails, 1)
				}
				chBuff <- 1
			}(&wg)
		} else {
			flag = false
		}
	}
	wg.Wait()
	if !flag {
		return ErrErrorsLimitExceeded
	}
	return nil
}
