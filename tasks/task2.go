package tasks

import (
	"fmt"
	"sync"
)

func Test2() bool {
	numbers := []int{2, 4, 6, 8, 10}

	// Инитим WG для синхронизации горутин
	var wg sync.WaitGroup
	
	// Добавляем ожидание равное количеству элементов массива
	wg.Add(len(numbers))

	for _, num := range numbers{
		go func(x int) {
			defer wg.Done()
			result := x * x
			fmt.Println(result)
		}(num)
	}

	// Дожидаемся горутин
	wg.Wait()

	return true
}