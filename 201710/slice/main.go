package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{}
	for i := 0; i < 4; i++ {
		start := time.Now()
		switch i {
		case 0:
			s = sample()
		case 1:
			s = sample1()
		case 2:
			s = sample2()
		case 3:
			s = sample3()
		default:
			break
		}
		fmt.Printf("%d: %5d ns ", i, time.Now().Sub(start).Nanoseconds())
		fmt.Println(s)
	}
}

func sample() (s []int) {
	s = []int{0, 1, 2, 3, 4}
	return
}

func sample1() (s []int) {
	s = make([]int, 0)
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
	return
}

func sample2() (s []int) {
	s = make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		s = append(s, i)
	}
	return
}

func sample3() (s []int) {
	s = make([]int, 5, 5)
	for i := 0; i < 5; i++ {
		s[i] = i
	}
	return
}
