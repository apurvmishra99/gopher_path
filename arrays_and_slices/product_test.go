package arrays_and_slices

import "testing"

func TestProduct(t *testing.T) {

	t.Run("calculate product of a collection", func(t *testing.T) {
		nums := []int{3, 6, 9}
		got := Product(nums)
		want := 162

		if got != want {
			t.Errorf("got %d but wanted %d, as given, %v", got, want, nums)
		}
	})

	t.Run("product of empty slice", func(t *testing.T) {
		nums := []int{}
		got := Product(nums)
		want := 0

		if got != want {
			t.Errorf("got %d but wanted %d, as given, %v", got, want, nums)
		}
	})
}
