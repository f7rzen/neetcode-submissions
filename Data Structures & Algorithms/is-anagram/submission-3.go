func isAnagram(s string, t string) bool {
	seenS := make(map[byte]int)
	seenT := make(map[byte]int)

	if len(s) != len(t) {
		return false
	}
	
	for i := 0; i < len(s); i++ {
		seenS[s[i]]++
		seenT[t[i]]++
	}

	for i := 0; i < len(s); i++ {
		if seenS[s[i]] != seenT[s[i]] {
			return false
		}
	}
	return true
}