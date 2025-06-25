package walnut_iq

func checkWalnutGenius(iqs []int) []int {
	result := make([]int, len(iqs))
	result[0] = iqs[0]
	for ind := 1; ind < len(iqs); ind++ {
		result[ind] = result[ind-1] + iqs[ind]
	}

	suffixSum := 0
	divNumber := len(iqs) - 1
	for ind := len(iqs) - 1; ind > 0; ind-- {
		result[ind] = (result[ind-1] + suffixSum) / divNumber
		suffixSum += iqs[ind]
	}
	result[0] = suffixSum / divNumber
	return result
}
