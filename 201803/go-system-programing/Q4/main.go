package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	chanLoop()
	contexts()
}

func chanLoop() {
	c := make(chan int)
	go func(c chan int) {
		defer close(c)
		for i := 0; i < 50; i++ {
			if i%7 == 0 {
				c <- i
			}
		}
	}(c)

	for i := range c {
		fmt.Println(i)
	}
}

func contexts() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go func() {
		defer cancel()
		time.Sleep(5 * time.Second)
	}()

	<-ctx.Done()
	fmt.Printf("done: %v", ctx.Err())
}
