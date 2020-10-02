package leetcode

func reverseString(s []byte) []byte {
	n := len(s)
	for i := 0; i < n/2; i++ {
		temp := s[i]
		s[i] = s[n-1-i]
		s[n-1-i] = temp
	}
	return s
}
