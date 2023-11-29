package tasks

import (
	"fmt"
	"sync"
)

func Test3() bool {
	numbers := []int{2, 4, 6, 8, 10}
	a := new(int)
	*a = 0
	// Инитим WG для синхронизации горутин
	var wg sync.WaitGroup
	
	// Добавляем ожидание равное количеству элементов массива
	wg.Add(len(numbers))

	for _, num := range numbers{
		go func(x int, sum *int) {
			defer wg.Done()
			*a += x * x
			//fmt.Println(result)
		}(num, a)
		fmt.Println(*a)
	}

	// Дожидаемся горутин
	wg.Wait()
	fmt.Println(*a)

	return true
}