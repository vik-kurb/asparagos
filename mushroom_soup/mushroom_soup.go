package mushroom_soup

func canPartition(mushrooms []int) bool {
	totalWeight := 0
	for _, mushroomWeight := range mushrooms {
		totalWeight += mushroomWeight
	}
	if totalWeight%2 == 1 {
		return false
	}
	potWeight := totalWeight / 2

	dpSumWeight := make([]bool, potWeight+1)
	dpSumWeight[0] = true
	for _, mushroomWeight := range mushrooms {
		for i := len(dpSumWeight) - 1; i >= 0; i-- {
			if dpSumWeight[i] && i+mushroomWeight < len(dpSumWeight) {
				dpSumWeight[i+mushroomWeight] = true
			}
		}
		if dpSumWeight[potWeight] {
			return true
		}
	}
	return false
}
