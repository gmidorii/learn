package main

import (
	"fmt"
	"time"
)

func main() {
	s := []int{}
	for i := 0; i < 5; i++ {
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
		case 4:
			s = sample4()
		default:
			break
		}
		fmt.Printf("%d: %5d ns ", i, time.Now().Sub(start).Nanoseconds())
		fmt.Println(s)
	}
}

func sample() []int {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return s
}

func sample1() []int {
	s := make([]int, 0)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	return s
}

func sample2() []int {
	s := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		s = append(s, i)
	}
	return s
}

func sample3() []int {
	s := make([]int, 10, 10)
	for i := 0; i < 10; i++ {
		s[i] = i
	}
	return s
}

func sample4() []int {
	s := make([]int, 10)
	for i := 0; i < 10; i++ {
		s[i] = i
	}
	return s
}
