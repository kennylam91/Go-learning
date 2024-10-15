package Variadic_spread_operator

func sum(nums ...int) int {
	var sum int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum

}
