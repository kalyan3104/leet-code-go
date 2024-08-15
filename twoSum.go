package main

import (
	"container/list"
	"fmt"
)

func twoSum(nums *list.List, target int) (int, int) {
	for e1 := nums.Front(); e1 != nil; e1 = e1.Next() {
		for e2 := e1.Next(); e2 != nil; e2 = e2.Next() {
			if e1.Value.(int)+e2.Value.(int) == target {
				return e1.Value.(int), e2.Value.(int)
			}
		}
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
	num1, num2 := twoSum(nums, target)
	fmt.Println(num1, num2) // Output: 2 7
}
