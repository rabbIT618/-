package main

import (
	"fmt"
)


func quickSort(nums []int) {  
    recursionSort(nums, 0, len(nums)-1)  
}  
  
func recursionSort(nums []int, left int, right int) {  
    if left < right {  
        pivot := partition(nums, left, right)  
        recursionSort(nums, left, pivot-1)  
        recursionSort(nums, pivot+1, right)  
    }  
}  
  
func partition(nums []int, left int, right int) int {  
    for left < right {  
        for left < right && nums[left] <= nums[right] {  
            right--  
        }  
        if left < right {  
            nums[left], nums[right] = nums[right], nums[left]  
            left++  
        }  
  
        for left < right && nums[left] <= nums[right] {  
            left++  
        }  
        if left < right {  
            nums[left], nums[right] = nums[right], nums[left]  
            right--  
        }  
    }  
    return left  
}

func main() {
	array := [...]int{25, 23, 1, 324, 65, 75, 30, 34}
	quickSort(array[:])
	for j := 0; j < len(array); j++ {
		fmt.Printf("%d\n", array[j])
	 }
}
