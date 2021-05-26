package main

import (
	"fmt"
	"github.com/ozoncp/ocp-tenant-api/internal/utils-lesson2"
	"github.com/ozoncp/ocp-tenant-api/internal/utils-lesson3"
)

func main() {
	lesson1()
	lesson2()
	lesson3()
}

func lesson1() {
	fmt.Println("Результаты первого урока")
	fmt.Println("Hi! This cmd will be done by Krivolutskiy Andrey.")
}

func lesson2() {
	fmt.Println("Результаты второго урока")
	lesson2TestSplitEmptyMass()
	lesson2TestSplitFullMass()
	lesson2SwapKeyValueEmpty()
	lesson2SwapKeyValueWithValue()
	lesson2RemoveElements()
}

func lesson2TestSplitEmptyMass() {
	utils_lesson2.Split(make([]int64, 0), 10)
}

func lesson2TestSplitFullMass() {
	a := make([]int64, 88)
	fmt.Println("Исходный массив:")
	for index, _ := range a {
		a[index] = int64(index)
		fmt.Printf(" %v", index)
	}
	fmt.Println()
	fmt.Println("Порезанный на кусочки массив:")
	b := utils_lesson2.Split(a, 10)
	for _, value := range b {
		for _, value2 := range value {
			fmt.Printf(" %v", value2)
		}
		fmt.Println()
	}
}

func lesson2SwapKeyValueEmpty() {
	utils_lesson2.SwapKeyValue(map[int64]int64{})
}

func lesson2SwapKeyValueWithValue() {
	baseMap := map[int64]int64{}
	baseMap[0] = 100
	baseMap[1] = 200
	baseMap[2] = 300
	fmt.Println("Исходный мапа:")
	for key, value := range baseMap {
		fmt.Printf("Key: %v; Value: %v \n", key, value)
	}
	swapResult := utils_lesson2.SwapKeyValue(baseMap)
	fmt.Println("Вывернутая мапа:")
	for key, value := range swapResult {
		fmt.Printf("Key: %v; Value: %v \n", key, value)
	}
}

func lesson2RemoveElements() {
	a := make([]int64, 28)
	fmt.Println("Исходный массив:")
	for index, _ := range a {
		a[index] = int64(index)
		fmt.Printf(" %v", index)
	}
	fmt.Println()
	fmt.Println("Надо исключить из массива:")
	b := utils_lesson2.Split(a, 10)
	for _, value2 := range b[1] {
		fmt.Printf(" %v", value2)
	}
	fmt.Println()
	fmt.Println("Результат:")
	removeResult := utils_lesson2.RemoveElements(a, b[1])
	for _, value := range removeResult {
		fmt.Printf(" %v", value)
	}
	fmt.Println()
}

func lesson3() {
	fmt.Println("Результаты третьего урока")
	createTenant()
}

func createTenant() {
	newTenant := utils_lesson3.Tenant{
		Id:   1,
		Name: "First",
		Type: 1,
	}
	fmt.Println("Создан первый елемент типа Tenant: " + utils_lesson3.ToString(newTenant))
}
