func groupAnagrams(strs []string) [][]string {

	seen := make(map[string][]string)

	for _, str := range strs {
		letters := strings.Split(str, "")
		sort.Strings(letters)
		sortedStr := strings.Join(letters, "")
		seen[sortedStr] = append(seen[sortedStr], str)
	}

	var result [][]string
	for _, group := range seen {
		result = append(result, group)
	}

	return result
}