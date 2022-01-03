package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	if m <= 0 {
		return nil
	}
	wg := sync.WaitGroup{}
	workerCount := n
	if len(tasks) < n {
		workerCount = len(tasks)
	}
	chTask := make(chan Task)
	chBuff := make(chan int, workerCount)
	for i := 0; i < workerCount; i++ {
		chBuff <- 0
	}
	go func() {
		var value int
		var failUntilExit int
		for _, task := range tasks {
			value = <-chBuff
			failUntilExit += value
			if failUntilExit < m {
				chTask <- task
			} else {
				close(chTask)
				return
			}
		}
		close(chTask)
	}()
	startedTaskCounter := 0
	for task := range chTask {
		startedTaskCounter++
		task := task
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := task()
			if err != nil {
				chBuff <- 1
			} else {
				chBuff <- 0
			}
		}()
	}
	wg.Wait()
	close(chBuff)
	var fails int
	for i := range chBuff {
		fails += i
	}
	if startedTaskCounter != len(tasks) || fails >= m {
		return ErrErrorsLimitExceeded
	}
	return nil
}
