package main

type ATable struct {
	best [][]string
	cur  [][]string
}

func NewATable(lenP, lenQ int) *ATable {
	A := new(ATable)
	A.best = make([][]string, lenP+1)
	A.cur = make([][]string, lenP+1)

	for i := 0; i < lenP; i++ {
		A.best[i] = make([]string, lenQ+1)
		A.cur[i] = make([]string, lenQ+1)
	}
	return A
}

func LCS(P, Q string) string {
	A := NewATable(len(P), len(Q))
	for i := 1; i <= len(P); i++ {
		for j := 1; j <= len(Q); j++ {
			if P[i-1] == Q[j-1] {
				A.cur[i][j] = A.cur[i-1][j-1] + string(Q[j-1])
			} else {
				A.cur[i][j] = ""
			}

			best := A.cur[i][j]
			if len(A.best[i-1][j]) > len(best) {
				best = A.best[i-1][j]
			}
			if len(A.best[i][j-1]) > len(best) {
				best = A.best[i][j-1]
			}
			if len(A.best[i-1][j-1]) > len(best) {
				best = A.best[i-1][j-1]
			}

			A.best[i][j] = best
		}
	}
	return A.best[len(P)][len(Q)]
}

func main() {

}
