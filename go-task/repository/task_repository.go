package repository

import "lobo.tech/task/model"

type TaskRepository struct {
}

func NewTaskRepository() *TaskRepository {

	return &TaskRepository{}

}

func (repository *TaskRepository) GetByID(id string) (model.Task, error) {

	var task model.Task = model.Task{id, "definir modulo", "", true}

	return task, nil

}
