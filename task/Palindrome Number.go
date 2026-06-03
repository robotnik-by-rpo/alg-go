package task

import (
	"math"
)

func IsPalindrome(x int) bool {
	if x < 0{
		return false
	}
	if x < 10 && x >= 0{
		return true
	}	
	len := int(math.Log10(float64(x)))+1
	numbers := make([]int,len)
	for i := 0; i < len ; i++{
		numbers[i] = x % 10
		x /= 10
	}
	for i, j := 0, len-1; i < j; i, j = i+1, j-1{
		if numbers[i] != numbers[j]{
			return false
		}
	}
	return true
}