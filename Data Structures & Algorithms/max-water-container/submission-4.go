func maxArea(heights []int) int {
	start := 0
	end := len(heights) - 1
	areaMax := (end - start) * min(heights[start], heights[end])
	for start < end {
		if heights[start] < heights[end] {
			start++
			area := (end - start) * min(heights[start], heights[end])
			if area > areaMax {
				areaMax = area
			}
		} else {
			end--
			area := (end - start) * min(heights[start], heights[end])
			if area > areaMax {
				areaMax = area
			}
		}
	}
	return areaMax
}