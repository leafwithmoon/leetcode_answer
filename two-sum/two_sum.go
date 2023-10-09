package twoSum

import "fmt"

func twoSum(nums []int, target int) []int {
	// store which num we have
	numMap := map[int]int{}

	for index, num := range nums {
		if value, exist := numMap[target-num]; exist {
			return []int{value, index}
		}
		numMap[num] = index
	}

	//  always have solution
	return []int{0, 0}
}

func Run() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))

}
