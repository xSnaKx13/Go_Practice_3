package main

import (
	"fmt"
)

func less(){
	arr := []int{1,2,3,4}
	printArr(&arr)
}

func printArr(array *[]int){
	arr := *array
	for i := len(*array) - 1; i >= 0; i--{
		fmt.Print(arr[i])
	}
}
