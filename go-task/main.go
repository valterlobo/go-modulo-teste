package main

import (
	"fmt"

	"github.com/go-playground/validator"
	"lobo.tech/task/repository"
)

func main() {
	taskRespository := repository.NewTaskRepository()
	task, err := taskRespository.GetByID("12233")

	if err == nil {
		fmt.Println(task.ID, task.Title, task.Done)

		v := validator.New()

		errField := v.Struct(task)

		if errField != nil {
			for _, e := range errField.(validator.ValidationErrors) {
				fmt.Println(e)
			}
		} else {

			fmt.Println("Tudo OK")

		}

	} else {
		fmt.Println(err)
	}

}
