package utils_lesson2

// Разделение на слайса на батчи - исходный слайс конвертировать в слайс слайсов - чанки одинкового размера (кроме последнего)
func Split(toSplit []int64, size int) [][]int64 {
	if len(toSplit) == 0 {
		return [][]int64{}
	}
	sizeOfBatches := len(toSplit) / size
	if len(toSplit)%size != 0 {
		sizeOfBatches += 1
	}
	batches := make([][]int64, sizeOfBatches, sizeOfBatches)
	index := 0
	for index+size < len(toSplit) {
		batches[index/size] = toSplit[index : index+size]
		index += size
	}
	batches[sizeOfBatches-1] = toSplit[index : index+len(toSplit)%size]
	return batches
}

// Обратный ключ - происходит конвертация отображения ("ключ-значение") в отображение ("значение-ключ")
func SwapKeyValue(baseMap map[int64]int64) map[int64]int64 {
	result := make(map[int64]int64, len(baseMap))
	for key, value := range baseMap {
		result[value] = key
	}
	return result
}

// Фильтрация по захордкоженному списку - нужно отфильтровать входной слайс по критерию отсутствия элемента в захардкоженном списке
// хардкодить не хочу - будет просто два слайса. Из одного выкидываем елементы второго
func RemoveElements(base []int64, toRemove []int64) []int64 {
	result := make([]int64, 0, len(base))
	for _, value := range base {
		elementFound := false
		for _, value2 := range toRemove {
			if value2 == value {
				elementFound = true
				break
			}
		}
		if !elementFound {
			result = append(result, value)
		}
	}
	return result
}
