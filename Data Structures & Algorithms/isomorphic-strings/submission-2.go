func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	isomerST := make(map[byte]byte)
	isomerTS := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		_, okST := isomerST[s[i]]
		if !okST {
			isomerST[s[i]] = t[i]
		} else {
			if isomerST[s[i]] != t[i] {
				return false
			}
		}

		_, okTS := isomerTS[t[i]]
		if !okTS {
			isomerTS[t[i]] = s[i]
		} else {
			if isomerTS[t[i]] != s[i] {
				return false
			}
		}
	}

	return true
}