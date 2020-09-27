package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var maxCandies int
	check := make([]bool, len(candies))

	for _, v := range candies {
		if v > maxCandies {
			maxCandies = v
		}
	}
	for i, v := range candies {
		if v+extraCandies >= maxCandies {
			check[i] = true
		}
	}

	return check
}
