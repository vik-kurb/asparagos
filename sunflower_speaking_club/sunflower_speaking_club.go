package sunflower_speaking_club

func findSpeakingPartner(englishLevel []int) []int {
	result := make([]int, len(englishLevel))
	stack := make([]int, 0, len(englishLevel))
	for ind := len(englishLevel) - 1; ind >= 0; ind-- {
		curLevel := englishLevel[ind]
		for len(stack) > 0 {
			stackInd := stack[len(stack)-1]
			if englishLevel[stackInd] > curLevel {
				break
			}
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			result[ind] = stack[len(stack)-1] - ind
		}
		stack = append(stack, ind)
	}
	return result
}
