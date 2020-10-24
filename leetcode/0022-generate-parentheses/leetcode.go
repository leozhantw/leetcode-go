package leetcode

func generateParenthesis(n int) []string {
	var (
		p   []byte
		res []string
	)
	gen(n, n, &p, &res)
	return res
}

func gen(ln, rn int, p *[]byte, res *[]string) {
	if ln == 0 && rn == 0 {
		*res = append(*res, string(*p))
		return
	}

	if ln > 0 {
		*p = append(*p, '(')
		gen(ln-1, rn, p, res)
		*p = (*p)[:len(*p)-1]
	}
	if rn > 0 && ln < rn {
		*p = append(*p, ')')
		gen(ln, rn-1, p, res)
		*p = (*p)[:len(*p)-1]
	}
}
