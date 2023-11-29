package tasks

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func CreateWorker(ctx context.Context, id int, wg *sync.WaitGroup, dataCh chan int) {
	defer wg.Done()
	

	for {
		select {
		case data := <-dataCh:
			fmt.Printf("Worker %d: Received data %d\n", id, data)
		case <-ctx.Done():
			fmt.Printf("Worker %d: Ended\n", id)
			return
		}
	}
}

func Test4() bool {
	dataCh := make(chan int)
	numWorkers := 5 // количество воркеров

	wg := &sync.WaitGroup{}
	wg.Add(numWorkers)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Запускаем воркеры
	for i := 1; i <= numWorkers; i++ {
		go CreateWorker(ctx, i, wg, dataCh)
	}




	// Записываем данные в канал (главный поток)
	go func(ctx context.Context) {
		for i := 1; ; i++ {
			select{
			case <-ctx.Done():
				return
			default:
				dataCh <- i
			}
		}
	}(ctx)

	// Ожидаем сигнала завершения программы
	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, os.Interrupt)
	<-stopCh
	cancel()

	go func() {
        <-stopCh
		cancel()
        close(dataCh)
    }()

	// Закрываем канал и ждем завершения всех воркеров
	
	wg.Wait()
	return true
}