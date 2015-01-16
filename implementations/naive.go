package implementations

type naiveType struct{}

var Naive = naiveType{}

func (p naiveType) Pascal(n, m int) int {
	rows := make([][]int, n+1)
	if n < 2 {
		rows = make([][]int, 2)
	}
	rows[0] = []int{1}
	rows[1] = []int{1, 1}
	for i := 2; i <= n; i++ {
		numColumns := i + 1
		rows[i] = make([]int, numColumns)
		rows[i][0] = 1
		rows[i][numColumns-1] = 1
		for j := 1; j < numColumns-1; j++ {
			rows[i][j] = rows[i-1][j-1] + rows[i-1][j]
		}
	}
	return rows[n][m]
}

func (p naiveType) String() string {
	return "Naive Implementation"
}
