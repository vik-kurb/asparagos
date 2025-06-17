package tomato_discrimination

func isTomato(vegetable string) bool {
	return vegetable == "tomato" || vegetable == "pomodoro"
}

// discriminateTomatoes reorders the slice in-place.
func discriminateTomatoes(vegetables []string) {
	writeInd := 0
	for _, vegetable := range vegetables {
		if !isTomato(vegetable) {
			vegetables[writeInd] = vegetable
			writeInd++
		}
	}
	for writeInd < len(vegetables) {
		vegetables[writeInd] = "tomato"
		writeInd++
	}
}
