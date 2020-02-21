package threesumclosest

import (
	"math"
	"sort"
)

func setResult(arr []int, result *int, sumSoFar int, target int) {

  if len(arr) == 0 {
    return
  }

	ind := sort.Search(len(arr), func(i int) bool {
		return target <= sumSoFar+arr[i]
	})

	if ind < 0 {
		ind++
	} else if ind >= len(arr) {
		ind--
	}

	if int(math.Abs(float64((sumSoFar+arr[ind])-target))) < int(math.Abs(float64((*result)-target))) {
		*result = sumSoFar + arr[ind]
	}
}

// ThreeSumClosest will return the closest sum to the target
func ThreeSumClosest(nums []int, target int) int {

	// sorting ints for lg(n) access
	sort.Ints(nums)

	result := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sumed := nums[i] + nums[j]
			setResult(nums[i+1:j], &result, sumed, target)
			setResult(nums[j+1:], &result, sumed, target)
		}
	}

	return result
}
