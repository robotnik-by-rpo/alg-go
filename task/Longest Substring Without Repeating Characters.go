package task

func LengthOfLongestSubstring(s string) int {
	last := make(map[rune]int)
	left := 0
	max_len := 0
	for right, c := range s{
		pos, exist := last[c]
		if exist && pos >= left{
			left = last[c] + 1
 		}
		last[c] = right
		max_len = max(max_len, right - left + 1)
	}
	return max_len
}