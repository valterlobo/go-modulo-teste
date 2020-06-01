package main

import (
	"fmt"

	anotacao_repo "lobo.tech/anotacao/repository"
	task_repo "lobo.tech/task/repository"
)

func main() {

	fmt.Println("APP")
	//ANOTACAO
	fmt.Println("ANOTACAO")
	anotacaoRespository := anotacao_repo.NewAnotacaoRepository()
	anotacao, err := anotacaoRespository.GetByID("3456")

	if err == nil {
		fmt.Println(anotacao.UUID, anotacao.Text)
	} else {
		fmt.Println(err)
	}

	//TASK
	fmt.Println("TASK")
	taskRespository := task_repo.NewTaskRepository()
	task, err := taskRespository.GetByID("12233")

	if err == nil {
		fmt.Println(task.ID, task.Title, task.Done)
	} else {
		fmt.Println(err)
	}

}
