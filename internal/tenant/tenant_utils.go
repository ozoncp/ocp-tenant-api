package tenant

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

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
	batches[butchCount-1] = tenant[index:]
	return batches
}

func ToMapWithId(input []Tenant) (map[uint64]Tenant, error) {
	result := make(map[uint64]Tenant, len(input))
	for _, value := range input {
		if mapElement, found := result[value.Id]; found {
			if mapElement == value {
				continue
			} else {
				return result, errors.New("Tenant::Id is not uniq")
			}
		} else {
			result[value.Id] = value
		}
	}
	return result, nil
}

func OpenFileInLoop() {
	fmt.Println("Работа с файлом:")
	fn := func(i int64) {
		file, err := os.Create("file.txt")
		check(err)
		n, errWrite := file.WriteString("i = " + strconv.FormatInt(i, 10))
		check(errWrite)
		fmt.Printf("Записано %d bytes\n", n)
		defer file.Close()
	}
	for i := int64(0); i < 5; i++ {
		fn(i)
	}
	file, err := os.Open("file.txt")
	check(err)
	defer file.Close()
	buffer := make([]byte, 20)
	size, errRead := file.Read(buffer)
	check(errRead)
	fmt.Println("Данные в файле:")
	fmt.Printf("%s\n", string(buffer[:size]))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
