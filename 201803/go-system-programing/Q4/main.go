package main

import "fmt"

func main() {
	chanLoop()
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
