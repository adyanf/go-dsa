package binary_search

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}

func BinarySearchRotated(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := low + ((high - low) / 2)

		if nums[mid] == target {
			return mid
		}

		if nums[low] <= nums[mid] {
			if nums[low] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else if target > nums[mid] && target <= nums[high] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
