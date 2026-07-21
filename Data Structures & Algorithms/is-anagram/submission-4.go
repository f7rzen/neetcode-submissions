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

	for ch, count := range seenS {
		if seenT[ch] != count {
			return false
		}
	}
	return true
}