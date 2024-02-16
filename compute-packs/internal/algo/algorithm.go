package algo

import (
	"sort"
)

// ComputePacking is a variation of the famous coin denomination problem.
// The difference is we don't give up when there is no exact solution,
// We find the nearest larger exact solution which should give us the minimum.
func ComputePacking(packs []int, order int) map[int]int {
	sort.Sort(sort.Reverse(sort.IntSlice(packs)))
	ceiling := (order/packs[0] + 1) * packs[0]

	upperBound := ceiling + 1
	dp := make([]int, upperBound)
	packChoice := make([]int, upperBound)
	for i := range dp {
		dp[i] = upperBound
	}

	dp[0] = 0
	for _, pack := range packs {
		for i := pack; i <= ceiling; i++ {
			if dp[i-pack] != upperBound && dp[i-pack]+1 < dp[i] {
				dp[i] = dp[i-pack] + 1
				packChoice[i] = pack
			}
		}
	}

	nearest := order
	for i := order; i < upperBound; i++ {
		if dp[i] != upperBound {
			nearest = i
			break
		}
	}

	solution := make(map[int]int, len(packs))
	for i := nearest; i > 0; i -= packChoice[i] {
		solution[packChoice[i]]++
	}

	return solution
}
