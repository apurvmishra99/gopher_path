package arrays_and_slices

// Sum : returns the sum of an integral array
func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	return sum
}

// SumAll : sum multiple arrays and return their sums as an array
func SumAll(numsToSum ...[]int) []int {
	lenOfNums := len(numsToSum)
	sums := make([]int, lenOfNums)

	for i, nums := range numsToSum {
		sums[i] = Sum(nums)
	}

	return sums
}

// SumAll : (alt) sum multiple arrays and return their sums as an array
func _SumAll(numsToSum ...[]int) []int {
	var sums []int

	for _, nums := range numsToSum {
		sums = append(sums, Sum(nums))
	}

	return sums
}

// SumAllTails : sum tails of multiple arrays and return their sums as an array
func SumAllTails(numsToSum ...[]int) (sums []int) {
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return
}
