package main

import (
	"fmt"
	"testing"
)

func BenchmarkInitDeclation(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		b.StopTimer()
		fmt.Sprint(s)
	}
}

func BenchmarkInitMakeLen0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0)
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkInitMakeLen10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 10)
		for i := 0; i < 10; i++ {
			s[i] = i
		}
	}
}

func BenchmarkInitMakeLen0Cap10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 0, 10)
		for i := 0; i < 10; i++ {
			s = append(s, i)
		}
	}
}

func BenchmarkInitMakeLen10Cap10(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := make([]int, 10, 10)
		for i := 0; i < 10; i++ {
			s[i] = i
		}
	}
}
