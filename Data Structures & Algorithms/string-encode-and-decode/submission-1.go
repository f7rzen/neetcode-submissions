type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var result string

	for _, str := range strs {
		result += strconv.Itoa(len(str)) + "#" + str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	var step string

	result := make([]string, 0)

	for i := 0; i < len(encoded); i++ {

		if encoded[i] != '#' {
			if _, err := strconv.Atoi(string(encoded[i])); err == nil {
				step += string(encoded[i])
			}
		} else {
			IntStep, _ := strconv.Atoi(step)
			result = append(result, encoded[i+1:i+1+IntStep])
			i += IntStep
			step = ""
		}
	}
	return result
}