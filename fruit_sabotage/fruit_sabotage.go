package fruit_sabotage

func findFruitKing(votes []int) int {
	var king int
	count := 0
	for _, vote := range votes {
		if count == 0 {
			king = vote
			count++
			continue
		}
		if king == vote {
			count++
		} else {
			count--
		}
	}
	return king
}
