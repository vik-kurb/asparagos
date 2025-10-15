package watermelon_football

func countWaysToMakeTeams(rating []int) int {
	oddSum := 0
	evenSum := 0
	for ind, r := range rating {
		if ind%2 == 0 {
			evenSum += r
		} else {
			oddSum += r
		}
	}
	var prefixOddSum, prefixEvenSum, count int
	for ind, r := range rating {
		newOddSum := prefixOddSum + (evenSum - prefixEvenSum)
		newEvenSum := prefixEvenSum + (oddSum - prefixOddSum)
		if ind%2 == 0 {
			newOddSum -= r
			prefixEvenSum += r
		} else {
			newEvenSum -= r
			prefixOddSum += r
		}
		if newOddSum == newEvenSum {
			count++
		}
	}
	return count
}
