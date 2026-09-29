
func maxArea(heights []int) int {
	start := 0
	end := len(heights) - 1
	areaMax := 0
	for start < end {
		area := (end - start) * min(heights[start], heights[end])

		if area > areaMax {
			areaMax = area
		}

		if heights[start] < heights[end] {
			start++
		} else {
			end--
		}
	}
	return areaMax
}