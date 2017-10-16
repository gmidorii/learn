package main

import (
	"fmt"
	"time"
)

func main() {
	sample()
	sample1()
}

func sample() {
	start := time.Now()
	s := []int{1, 2, 3, 4, 5}
	s = append(s, 6)
	fmt.Println(s)
	fmt.Printf("%d ns\n", time.Now().Sub(start).Nanoseconds())
}

func sample1() {
	start := time.Now()
	s := make([]int, 0)
	for i := 1; i < 6; i++ {
		s = append(s, i)
	}
	s = append(s, 6)
	fmt.Println(s)
	fmt.Printf("%d ns\n", time.Now().Sub(start).Nanoseconds())
}
