package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func doWorkWithValue(ctx context.Context) {
	output := fmt.Sprintf("Doing some work with Value: %s", ctx.Value("myKey"))
	fmt.Println(output)
}

func doWorkWithTimeout(ctx context.Context, workChan <-chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("Something went wrong: %s\n", err)
				wg.Done()
			}
			fmt.Printf("worker: finished\n")
		case i := <-workChan:
			fmt.Printf("squaring int: %d -> %d\n", i, i*i)
			wg.Done()
		}
	}
}

func main() {
	ctx := context.WithValue(context.TODO(), "myKey", "Hello World")

	doWorkWithValue(ctx)

	workChan := make(chan int, 1)

	ctxTwo, cancel := context.WithTimeout(context.TODO(), time.Second*3)
	defer cancel()

	wg := &sync.WaitGroup{}

	fmt.Println("starting worker...")

	go doWorkWithTimeout(ctxTwo, workChan, wg)

	ints := []int{2, 4, 6, 4, 11}

	wg.Add(1)

	go func() {
		for i := range ints {
			wg.Add(1)
			workChan <- i
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
}
