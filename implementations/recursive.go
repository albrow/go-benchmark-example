package implementations

type recursiveType struct{}

var Recursive = &recursiveType{}

func (p *recursiveType) Pascal(n, m int) int {
	if n == 0 || m == 0 || n == m {
		return 1
	}
	return p.Pascal(n-1, m-1) + p.Pascal(n-1, m)
}

func (p *recursiveType) String() string {
	return "Recursive Implementation"
}
