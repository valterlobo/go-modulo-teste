package main

import (
	"fmt"

	"lobo.tech/task/repository"
)

func main() {
	taskRespository := repository.NewTaskRepository()
	task, err := taskRespository.GetByID("12233")

	if err == nil {
		fmt.Println(task.ID, task.Title, task.Done)
	} else {
		fmt.Println(err)
	}

}
