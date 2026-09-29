func isPalindrome(s string) bool {
	start := 0
	s = strings.ToUpper(s)
	runes := []rune(s)
	end := len(runes) - 1

	for start < end {
		for start < end {
			if !unicode.IsLetter(runes[start]) && !unicode.IsDigit(runes[start]) {
				start++
			} else {
				break
			}
		}

		for start < end {
			if !unicode.IsLetter(runes[end]) && !unicode.IsDigit(runes[end]) {
				end--
			} else {
				break
			}
		}
		if runes[start] != runes[end] {
			return false
		}
		start++
		end--
	}
	return true
}