package task


func MaxArea(height []int) int {
	maxArea := 0
	left := 0
	lenArr := len(height)
	right := lenArr -1
	for left < right {
		num := (right - left)*min(height[left], height[right])
		maxArea = max(maxArea, num)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}