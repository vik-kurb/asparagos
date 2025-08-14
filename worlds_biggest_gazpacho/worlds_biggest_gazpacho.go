package worlds_biggest_gazpacho

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func collectTomatoes(beds []int) int {
	takeSum := 0
	notTakeSum := 0
	for _, num := range beds {
		takeSum, notTakeSum = notTakeSum+num, max(notTakeSum, takeSum)
	}
	return max(takeSum, notTakeSum)
}
