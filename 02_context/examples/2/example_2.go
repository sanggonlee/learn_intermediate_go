package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	_ = cancel // Discard the cancel function, it's not used in this example.

	go func() {
		<-ctx.Done()
		fmt.Println("Context done!")
	}()

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
