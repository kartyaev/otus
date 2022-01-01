package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	chTask := make(chan Task)
	chBuff := make(chan int, n)
	for i := 0; i < n; i++ {
		chBuff <- 0
	}
	go func() {
		var fails int
		var value int
		for _, task := range tasks {
			value = <-chBuff
			fails = fails + value
			if fails < m {
				chTask <- task
			} else {
				close(chTask)
				return
			}
		}
		close(chTask)
	}()
	littleCounter := 0
	for task := range chTask {
		littleCounter++
		task := task
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			err := task()
			if err != nil {
				chBuff <- 1
			} else {
				chBuff <- 0
			}
		}(&wg)
	}
	wg.Wait()
	if littleCounter != len(tasks) {
		return ErrErrorsLimitExceeded
	}
	return nil
}
