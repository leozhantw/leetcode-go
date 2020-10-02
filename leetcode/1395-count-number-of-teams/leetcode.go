package leetcode

func numTeams(rating []int) int {
	var teams int
	n := len(rating)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				if (rating[i] < rating[j] && rating[j] < rating[k]) || (rating[i] > rating[j] && rating[j] > rating[k]) {
					teams++
				}
			}
		}
	}
	return teams
}
