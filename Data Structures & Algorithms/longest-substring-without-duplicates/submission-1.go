func lengthOfLongestSubstring(s string) int {
	left := 0
	maxLen := 0
	sl := make(map[byte]bool)
	for right := 0; right < len(s); {
		_, ok := sl[s[right]]
		if !ok {
			sl[s[right]] = true

			mLen := len(s[left : right+1])
			if mLen > maxLen {
				maxLen = mLen
			}

			right++
		} else {
			delete(sl, s[left])
			left++
		}
	}
	return maxLen
}