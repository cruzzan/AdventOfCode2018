package one

import "fmt"

func Run(task string)  {
	if len(task) != 0 {
		fmt.Println("Running specific task")
	} else {
		fmt.Println("Running full day")
		task1()
	}
}

func task1()  {
	fmt.Println("Task 1")
}
