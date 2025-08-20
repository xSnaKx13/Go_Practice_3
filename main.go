package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Введите неограниченное кол-во чисел через запятую:  ")
	numberStr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	numberStr = strings.TrimSpace(numberStr)
	numberString := strings.Split(numberStr, ",")

	var numbersInt []int

	for _, elem := range numberString {
		numStr := strings.TrimSpace(elem)
		num, _ := strconv.Atoi(numStr)
		numbersInt = append(numbersInt, num)
	}

	var countIdx int

	var sum int
	var avg int

	for _, elem := range numbersInt {
		sum += elem
		countIdx++
	}
	avg = sum / countIdx
	sort.Ints(numbersInt)
	lenght := len(numbersInt)
	var median float64

	if lenght%2 == 0 {
		median = float64(numbersInt[lenght/2-1]) + float64(numbersInt[lenght/2])/2
	} else {
		median = float64(numbersInt[lenght/2])
	}

	fmt.Println(numbersInt)
	fmt.Printf("Сумма чисел равна: %d\n", sum)
	fmt.Printf("Среднее значение равно: %d\n", avg)
	fmt.Printf("Медиана чисел равна: %.2f\n", median)
	less()
}
