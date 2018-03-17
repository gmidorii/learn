package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ChanLoop()
	Contexts()
	//Signals()
}

// ChanLoop is sample code combined by goroutine, channel and for-loop.
func ChanLoop() {
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

// Contexts is sample code implemented context timeout.
func Contexts() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	go func() {
		defer cancel()
		time.Sleep(5 * time.Second)
	}()

	<-ctx.Done()
	fmt.Printf("done: %v\n", ctx.Err())
}

// Signals is sample code signal notify to channel.
func Signals() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	fmt.Println("Waiting ...")
	<-sig
	fmt.Println("SIGINT arrive")
}
