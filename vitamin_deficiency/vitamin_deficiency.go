package vitamin_deficiency

func findVitaminsSum(vitamins []int, targetSum int) []int {
	vitaminToLastInd := make(map[int]int)
	for ind, vitamin := range vitamins {
		targetValue := targetSum - vitamin
		if targetInd, ok := vitaminToLastInd[targetValue]; ok {
			return []int{targetInd, ind}
		}
		vitaminToLastInd[vitamin] = ind
	}
	return nil
}
