package orange_seeds

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sowTheSeedsOfDiscord(sectors []int) int {
	maxInd := 0
	nextMaxInd := 0
	throws := 0
	for ind, distance := range sectors {
		if maxInd >= len(sectors)-1 {
			return throws
		}
		if ind > maxInd {
			maxInd = nextMaxInd
			nextMaxInd = ind + distance
			throws++
		} else {
			nextMaxInd = max(nextMaxInd, ind+distance)
		}
	}
	return throws
}
