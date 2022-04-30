package main

import (
	"fmt"
	"time"
	executor "github.com/Braurbeki/multithreading_lab2/executors"
)

func main() {
	const max_workers = 2
	exec := executor.CustomExecutor{max_workers, 4}
	nums := [4]int{1, 2, 3, 4}
	future := exec.Custom_map(longRunningTask, nums[:])
	for i := 0; i < future.Size(); i++ {
		if i == 0 || i == 2 {
			fmt.Printf("Time: %s\n", time.Now().String())
		}
		
		fmt.Printf("Value: %d\n", future.Result())
	}
}

func longRunningTask(x int) int {
	time.Sleep(time.Second * 2)
	return x * 2
}