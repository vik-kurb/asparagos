package twin_salads

import (
	"strings"
)

func areSaladsTwin(vegetables []string) bool {
	if len(vegetables) < 2 {
		return true
	}
	leftInd := 0
	rightInd := len(vegetables) - 1
	for leftInd < rightInd {
		if strings.ToLower(vegetables[leftInd]) != strings.ToLower(vegetables[rightInd]) {
			return false
		}
		leftInd++
		rightInd--
	}
	return true
}
