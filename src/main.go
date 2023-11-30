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
		
		case "-5":
			a := tasks.Test5()
			if a {
				log.Println("Task 5 passed")
				continue
			}
			log.Println("Problem with task 5")
			continue
		
		case "-6":
			a := tasks.Test6()
			if a {
				log.Println("Task 6 passed")
				continue
			}
			log.Println("Problem with task 6")
			continue
		
		case "-7":
			a := tasks.Test7()
			if a {
				log.Println("Task 7 passed")
				continue
			}
			log.Println("Problem with task 7")
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
