package executors

import (
	// "fmt"
	"sync"
)

var queue chan int
var WG sync.WaitGroup

type CustomExecutor struct {
	Max_workers int
	Size int
}

type workerThread struct {
	foo func(int) int
}

type Future struct {
	val chan int
	siz int
}

func (e CustomExecutor) Custom_map(foo func(int) int, nums []int) *Future {
	queue = make(chan int, e.Max_workers)
	res := Future{make(chan int, e.Size), e.Size}
	go e.fill_queue(nums[:])
	go e.notify_workers(foo, &res)
	return &res
}

func (e CustomExecutor) fill_queue(nums []int) {
	for i := 0; i < e.Size; i++ {
		queue <- nums[i]
	}
}

func (e CustomExecutor) notify_workers(foo func(int) int, res *Future) {
	w := workerThread{foo}
	for i := 0; i < e.Size; i++ {
		if i != 0 && i % e.Max_workers - 1 == 0 {
			WG.Add(e.Max_workers)
			// fmt.Printf("WG added i: %d. Queue len: %d\n", i, len(queue))
		}
		// fmt.Printf("Worker started. i: %d\n", i)
		go w.run(res)
		WG.Wait()
	}
	
}

func (w workerThread) run(f* Future) {
	defer WG.Done()
	f.val <- w.foo(<-queue)
	// fmt.Println("Worker finished")
}

func (f *Future) Result() int {
	return <-f.val
}

func (f *Future) Size() int {
	return f.siz
}