package arrays_and_slices

// Product : returns the product of an array safely
func Product(nums []int) int {
	prod := 1
	if len(nums) == 0 {
		return 0
	}
	for _, num := range nums {
		prod *= num
	}

	return prod
}
