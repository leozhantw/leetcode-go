package leetcode

func processQueries(queries []int, m int) []int {
	P := make([]int, m)
	for i := 0; i < m; i++ {
		P[i] = i + 1
	}

	result := make([]int, len(queries))
	for k, v := range queries {
		for i := 0; i < m; i++ {
			if P[i] == v {
				copy(P[1:i+1], P[0:i])
				P[0] = v
				result[k] = i
			}
		}
	}
	return result
}
