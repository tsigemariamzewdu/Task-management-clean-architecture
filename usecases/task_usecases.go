package usecases

import (
	"context"
	"errors"
	domain "task_management/Domain"
	repositories "task_management/Repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//define TaskUseCase struct
type TaskUseCase struct {
	TaskRepo repositories.TaskRepository
}

//add new task usecase
func (uc *TaskUseCase) AddTask(ctx context.Context,input *domain.InputTask)(*domain.Task,error){
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	task := &domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       input.Title,
		Description: input.Description,
		Status:   input.Status,
		
	}

	err:= uc.TaskRepo.CreateTask(ctx,task)
	if err!=nil{
		return nil,errors.New("failed to create task")
	}
	return task,nil
}
//getalltasksusecase
func (uc *TaskUseCase) GetAllTasks(ctx context.Context)([]domain.Task ,error){
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	tasks,err:=uc.TaskRepo.GetAllTasks(ctx)
	if err !=nil{
		return nil,errors.New("failed to retrieve")
	}
	return tasks,nil

}
//get task byID use case
func (uc *TaskUseCase) GetTaskByID(ctx context.Context,id string)(*domain.Task,error){
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	
	
	task,err:=uc.TaskRepo.GetTaskByID(ctx,id)
	if err!=nil{
		return nil,errors.New("task not found")
	}
	return task,nil
}
//update task by id
func (uc *TaskUseCase) UpdateTaskByID(ctx context.Context, id string, input *domain.Task) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()


	return uc.TaskRepo.UpdateTaskByID(ctx, id, input)
}
//delete task by id
func (uc *TaskUseCase) DeleteTaskByID(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	return uc.TaskRepo.DeleteTaskByID(ctx,id)
}