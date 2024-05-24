package fast_and_slow_pointers

import "go-dsa/structs"

func IsHappy(num int) bool {
	sp, fp := num, sumOfSquaredDigits(num)

	for sp != fp && fp != 1 {
		sp = sumOfSquaredDigits(sp)
		fp = sumOfSquaredDigits(sumOfSquaredDigits(fp))
	}
	if fp == 1 {
		return true
	}
	return false
}

func DetectCycle(linkedList *structs.LinkedList) bool {
	if linkedList == nil || linkedList.Head == nil {
		return false
	}
	head := linkedList.Head

	sp, fp := head, head.Next
	for sp != nil && fp != nil && sp != fp {
		sp = sp.Next
		if fp.Next != nil {
			fp = fp.Next.Next
		} else {
			fp = fp.Next
		}
	}

	if sp == nil || fp == nil {
		return false
	}

	return true
}

func sumOfSquaredDigits(num int) int {
	sum := 0
	cur := num
	for cur != 0 {
		digit := cur % 10
		sum += digit * digit
		cur = cur / 10
	}

	return sum
}
