package main

import (
	"fmt"
)

func less(){
	arr := []int{1,2,3,4}
	printArr(&arr)
}

func printArr(*array[]int){
	for (int i = len(array), i > 0, i--){
		fmt.Print(arr[i])
	}
}
