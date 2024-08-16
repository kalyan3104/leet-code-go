package main

import (
	"container/list"
	"fmt"
)

func twoSum(nums *list.List, target int) (int, int) {
	// Create a map to store the values and their corresponding indices
	valueToIndex := make(map[int]int)
	
	// Iterate through the list with a counter to track indices
	i := 0
	for e := nums.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)
		complement := target - val
		
		if index, found := valueToIndex[complement]; found {
			return index, i
		}
		
		valueToIndex[val] = i
		i++
	}
	
	return -1, -1
}

func main() {
	nums := list.New()
	nums.PushBack(2)
	nums.PushBack(7)
	nums.PushBack(11)
	nums.PushBack(15)

	target := 9
	i, j := twoSum(nums, target)
	fmt.Println(i, j) // Output: 0 1
}
