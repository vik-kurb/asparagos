package peanuts_ego

func findSmartestPeanut(peanuts []int, k int) []int {
	res := make([]int, 0, len(peanuts)-k+1)
	deque := make([]int, 0, len(peanuts))
	for i, peanutLevel := range peanuts {
		for len(deque) > 0 && peanuts[deque[len(deque)-1]] <= peanutLevel {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, i)

		if len(deque) > 1 && deque[0] <= i-k {
			deque = deque[1:]
		}
		if i+1 >= k {
			res = append(res, peanuts[deque[0]])
		}
	}
	return res
}
