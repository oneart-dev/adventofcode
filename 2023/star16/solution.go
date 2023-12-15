package star16

import (
	"fmt"
	"strings"
)

func Fun(input []string) interface{} {
	result := 0

	for _, line := range input {

		if len(line) == 0 {
			continue
		}

		nums := strings.Split(line, " ")
		if len(nums) < 2 {
			panic("Invalid input")
		}

		num := make([]int, len(nums))
		for i, n := range nums {
			fmt.Sscanf(n, "%d", &num[i])
		}

		result += calc(num)
	}

	return result
}

func calc(nums []int) int {
	subNums := make([]int, 0, len(nums)-1)
	allZero := true
	for i := range nums {
		if i < 1 {
			continue
		}

		subNums = append(subNums, nums[i]-nums[i-1])
		if nums[i-1] != 0 {
			allZero = false
		}
	}

	if allZero {
		return 0
	}

	r := calc(subNums)
	return nums[len(nums)-1] + r
}
