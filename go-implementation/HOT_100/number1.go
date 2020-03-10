package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	fmt.Println(twoSum(nums, 9))
	fmt.Println(twoSum1(nums, 9))
}

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums{
		another_num := target -  v
		if _, ok := hashmap[another_num]; ok {
			return []int{hashmap[another_num], i}
		}
		hashmap[v] = i
	}
	return nil
}

func  twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target - v]; ok && i != j {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}