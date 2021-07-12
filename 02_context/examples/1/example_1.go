package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	time.AfterFunc(4*time.Second, func() {
		cancel()
	})

	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("After 1 second:", ctx.Err())
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("After 3 seconds:", ctx.Err())
	}()

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("After 5 seconds:", ctx.Err())
	}()

	time.Sleep(6 * time.Second)
}
