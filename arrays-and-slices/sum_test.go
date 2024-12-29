package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("colection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d, give %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum of each slice n of a set of n slices", func(t *testing.T) {
		var got []int = SumAll([]int{1, 2}, []int{1, 2, 3})
		var want []int = []int{3, 6}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(slice1, slice2 []int, t *testing.T) {
		if !reflect.DeepEqual(slice1, slice2) {
			t.Errorf("got %v want %v", slice1, slice2)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{1, 9, 1})
		want := []int{2, 10}
		checkSums(got, want, t)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 9, 1})
		want := []int{0, 10}
		checkSums(got, want, t)
	})
}
