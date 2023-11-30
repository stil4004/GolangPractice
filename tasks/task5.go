package tasks

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Worker struct {
	wg  *sync.WaitGroup
	ch  chan int
	ctx context.Context
}

func createSender(wrk *Worker){
	defer wrk.wg.Done()
	for{
		select{
		case <- wrk.ctx.Done(): 
			fmt.Printf("stop sender\n")
			return
		default:
			a := rand.Intn(50)
			wrk.ch <- a
			fmt.Printf("sended: %d\n", a)
		}
	}
}

func createReceiver(wrk *Worker){
	defer wrk.wg.Done()
	for{
		select{
		case num := <-wrk.ch:
			fmt.Printf("received: %d\n", num)
		case <-wrk.ctx.Done():
			fmt.Println("stop receiver\n")
			return
		}
	}
}

func Test5() bool {

	dataCh := make(chan int)

	rand.Seed(404)

	var n int

	_, err := fmt.Scan(&n)
	if err != nil{
		log.Fatalf("couldn't scan data: %v\n", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	worker_sender := &Worker{
		wg: wg,
		ch: dataCh,
		ctx: ctx,
	}

	worker_receiver := &Worker{
		wg: wg,
		ch: dataCh,
		ctx: ctx,
	}

	go createSender(worker_sender)
	go createReceiver(worker_receiver)

	go func(n_time int) {
		time.Sleep(time.Duration(n_time) * time.Second)
		cancel()
        close(dataCh)
    }(n)

	wg.Wait()

	return true
}
