package utils_lesson3

import (
	"errors"
	"strconv"
)

type Tenant struct {
	Id   uint64
	Name string
	Type uint8
}

func ToString(input Tenant) string {
	return "Id: " + strconv.FormatUint(input.Id, 5) +
		"; Name: " + input.Name +
		"; Type: " + strconv.FormatUint(uint64(input.Type), 2)
}

func SplitToButch(tenant []Tenant, butchSize uint) [][]Tenant {
	if len(tenant) == 0 {
		return [][]Tenant{}
	}
	lenTenant := uint(len(tenant))
	butchCount := lenTenant / butchSize
	if lenTenant%butchSize != 0 {
		butchCount += 1
	}
	batches := make([][]Tenant, butchCount)
	index := uint(0)
	for index+butchSize < lenTenant {
		batches[index/butchSize] = tenant[index : index+butchSize]
		index += butchSize
	}
	batches[butchCount-1] = tenant[index : index+lenTenant%butchSize]
	return batches
}

func ToMapWithId(input []Tenant) (map[uint64]Tenant, error) {
	result := make(map[uint64]Tenant, len(input))
	myError := errors.New("OK")
	for _, value := range input {
		if mapElement, found := result[value.Id]; found {
			if mapElement == value {
				continue
			} else {
				myError = errors.New("Tenant::Id is not uniq")
				break
			}
		} else {
			result[value.Id] = value
		}
	}
	return result, myError
}

//Используя defer и функтор реализовать открытие и закрытие файла в цикле
