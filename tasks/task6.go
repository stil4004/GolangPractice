package tasks

import (
	"context"
	"fmt"
	"time"
)

func Test6() bool {
	// Способы осознанной оставновки горутины
	// 1 - контекст
	// 2 - каналы
	a := testContext();
	if a {
		fmt.Println("test context passed")
	}
	b := testChannel(); 
	if b {
		fmt.Println("test Chanel passed")
	}

	return a && b
}

func testContext() bool {
	
	// Создаем канао для ответа
	ansCh := make(chan bool)

	// Создаем контекст по которому вызываем откл
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// Запускаем горутину в которой:
	// Либо вернется по контексту (тогда все ок)
	// Либо что все выполнится (тогда все в контексте задачи - плохо)
	go func(ctx context.Context, ansChan chan bool){
		for{
			select{
			case <-time.After(10 * time.Second):
				fmt.Println("All work done (context version)")
				ansChan<-false
			case <- ctx.Done():
				fmt.Println("Ended by context")
				ansChan<-true
				return
			}
		}

	}(ctx, ansCh)

	// Вызовем отруб через 3 секунды
	go func(){
		time.Sleep(3 * time.Second)
		cancel()
	}()

	ans := <-ansCh
	close(ansCh)

	return ans
}

func doWork() int{
	time.Sleep(10 * time.Second)
	return 13
}

func testChannel() bool {
	stopCh := make(chan bool)
	ch := make(chan int)

	go func(){
		for{
			select{
			case ch <- doWork():
			case <-stopCh:
				close(ch)
				fmt.Println("Ended by chan")
				return
			}
		}
	}()

	// Вызовем отруб через 3 секунды
	go func(){
		time.Sleep(3 * time.Second)
		stopCh <- true
	}()
	
	for i := range ch {
        fmt.Println("receive value: ", i)
    }
	
	close(stopCh)

	return true
}