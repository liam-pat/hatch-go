package main

import "fmt"

/**
Example:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].

*/

// o(n) = n*n
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// o(n) = n
func twoSum01(nums []int, target int) []int {
	map01 := map[int]int{}
	for sIndex, value := range nums {
		if mIndex, ok := map01[target-value]; ok {
			return []int{mIndex, sIndex}
		}
		map01[value] = sIndex
	}
	return nil
}

func test(nums []int, target int) []int {

	mapSaver := make(map[int]int)

	for key, value := range nums {
		if _, ok := mapSaver[target-value]; ok {

			return []int{mapSaver[target-value], key}
		}

		mapSaver[value] = key
	}

	return nil
}

func main() {

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))

	fmt.Println(twoSum01([]int{3, 2, 4}, 6))

	fmt.Println(test([]int{3, 2, 4}, 6))

}
