/**
题目

两数之和
*/
package no001

func twoSum(nums []int, target int) []int {
	cacheMap := make(map[int]int, len(nums))
	for index, value := range nums {
		if v, ok := cacheMap[target-value]; ok {
			return []int{v, index}
		}
		cacheMap[value] = index
	}
	return nil
}
