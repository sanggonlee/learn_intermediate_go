package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sanggonlee/asyncutil"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	for err := range asyncutil.Collect(
		fnThatReturnsErrorAfterNSeconds(ctx, 1, nil),
		fnThatReturnsErrorAfterNSeconds(ctx, 2, errors.New("err")),
		fnThatReturnsErrorAfterNSeconds(ctx, 3, nil),
	) {
		if err != nil {
			fmt.Println("Error:", err)
			cancel()
		}
	}
}

func fnThatReturnsErrorAfterNSeconds(ctx context.Context, n int, err error) chan error {
	errs := make(chan error)
	go func() {
		defer close(errs)

		select {
		case <-ctx.Done():
			errs <- ctx.Err()
		case <-time.After(time.Duration(n) * time.Second):
			errs <- err
		}
	}()
	return errs
}
