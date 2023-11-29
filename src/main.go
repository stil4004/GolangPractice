package main

import (
	"GolangPractice/tasks"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args{
		switch arg{
		case "-1":
			a := tasks.Test1()
			if a {
				log.Println("Task 1 passed")
				continue
			}
			log.Println("Problem with task 1")
			continue
		case "-2":
			a := tasks.Test2()
			if a {
				log.Println("Task 2 passed")
				continue
			}
			log.Println("Problem with task 2")
			continue	
		case "-3":
			a := tasks.Test3()
			if a {
				log.Println("Task 3 passed")
				continue
			}
			log.Println("Problem with task 3")
			continue
		
		case "-4":
			a := tasks.Test4()
			if a {
				log.Println("Task 4 passed")
				continue
			}
			log.Println("Problem with task 4")
			continue
		
		default:
			RunAllTests()
			break
		}
	}
	fmt.Println("Bye...")
}

func RunAllTests() bool{

	// Task 1
	right := tasks.Test1()
	if !right{
		return false
	}
	log.Println("Task 1 passed")

	// Task 2
	right = tasks.Test2()
	if !right{
		return false
	}
	log.Println("Task 2 passed")


	return true
}
