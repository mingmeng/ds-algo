package arraylist

func findPoisonedDuration(timeSeries []int, duration int) int {
	length := len(timeSeries)
	if length == 0 {
		return 0
	}

	result := duration
	for i := length - 1; i > 0; i-- {
		subs := timeSeries[i] - timeSeries[i-1]
		if subs > duration {
			result += subs
		} else {
			result += subs
		}
	}

	return result
}
