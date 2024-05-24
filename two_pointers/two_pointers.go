package two_pointers

import "sort"

func IsPalindrome(inputString string) bool {
	left := 0
	right := len(inputString) - 1
	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func SumOfThreeValues(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-2; i++ {
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if sum == target {
				return true
			} else if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}

	return false
}
