package main

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of variable length", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		got := Sum(numbers)
		want := 55

		if got != want {
			t.Errorf("want '%d', got '%d', given: %v", want, got, numbers)
		}
	})
}

func ExampleSum() {
	numbers := []int{2, 4, 6}
	sum := Sum(numbers)
	fmt.Println(sum)
	// Output: 12
}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}

}
