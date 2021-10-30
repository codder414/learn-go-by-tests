package main

import (
	"fmt"
	"reflect"
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

func TestSumAll(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want '%d' but got '%d'", want, got)
		}
	}

	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 5, 6})
		want := []int{5, 11}
		checkSums(t, got, want)
	})

	t.Run("make sunm of empty slice", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		checkSums(t, got, want)
	})
}

func ExampleSumAllTails() {
	numbersA := []int{1, 2, 3}
	numbersB := []int{4, 5, 6}

	sums := SumAllTails(numbersA, numbersB)

	fmt.Printf("%v", sums)

	// Output: [5 11]
}

func BenchmarkSumAllTails(b *testing.B) {
	numbersA := []int{1, 2, 3}
	numbersB := []int{4, 5, 6}
	for i := 0; i < b.N; i++ {
		SumAllTails(numbersA, numbersB)
	}
}
