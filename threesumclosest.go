package threesumclosest

import (
	"math"
	"sort"
)

// ThreeSumClosest will return the closest sum to the target
func ThreeSumClosest(nums []int, target int) int {

	// sorting ints for lg(n) access
	sort.Ints(nums)

	result := math.MaxInt32

	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

		}
	}
}
