package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// 获取一个空的跟context
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	go func() {
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {

	select {
	case <-time.After(5 * time.Second): // 5 seconds pass
		fmt.Println("finish doing something")
	case <-ctx.Done(): // ctx is cancelled
		err := ctx.Err()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	time.Sleep(5 * time.Second)
	fmt.Println("finish doing something...")
}
