package leetcode

func numTeams(rating []int) int {
	var teams int
	n := len(rating)
	for i := 1; i < n-1; i++ {
		var leftSmall, leftLarge, rightSmall, rightLarge int
		for j := i - 1; j >= 0; j-- {
			if rating[j] < rating[i] {
				leftSmall++
			} else if rating[j] > rating[i] {
				leftLarge++
			}
		}
		for j := i + 1; j < n; j++ {
			if rating[j] < rating[i] {
				rightSmall++
			} else if rating[j] > rating[i] {
				rightLarge++
			}
		}
		teams += (leftSmall * rightLarge) + (leftLarge * rightSmall)
	}
	return teams
}
