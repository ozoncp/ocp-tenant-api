package main

import (
	"fmt"
	"github.com/ozoncp/ocp-tenant-api/internal/utils"
)

func main() {
	fmt.Println("Hi! This cmd will be done by Krivolutskiy Andrey.")

	// делим массив
	a := make([]int64, 88)
	for index, _ := range a {
		a[index] = int64(index)
		fmt.Printf(" %v", index)
	}
	fmt.Println()
	fmt.Println()
	b := utils.Split(a, 10)
	for _, value := range b {
		for _, value2 := range value {
			fmt.Printf(" %v", value2)
		}
		fmt.Println()
	}
	// выворачиваем мапу
	baseMap := map[int64]int64{}
	baseMap[0] = 100
	baseMap[1] = 200
	baseMap[2] = 300
	swapResult := utils.SwapKeyValue(baseMap)
	for key, value := range swapResult {
		fmt.Printf("%v %v \n", key, value)
	}
	// удаляем первые 10 елементов
	removeResult := utils.RemoveElements(a, b[0])
	for _, value := range removeResult {
		fmt.Printf(" %v", value)
	}
}
