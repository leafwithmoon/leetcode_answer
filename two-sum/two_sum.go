package TwoSum

import "fmt"

func twoSum(nums []int, target int) []int {
	// store which num we have
	numMap := map[int]bool{}

	for _, num := range nums {
		fmt.Println(num)
		if numMap[target-num] {
			return []int{num, target - num}
		}
		numMap[num] = true
	}

	//  always have solution
	return []int{0, 0}
}

func TwoSumRun() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
