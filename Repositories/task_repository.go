package repositories

import (
	"context"
	"errors"
	"fmt"
	domain "task_management/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepository interface {
	CreateTask(ctx context.Context, task *domain.Task) error
	GetAllTasks(ctx context.Context) ([]domain.Task, error)
	GetTaskByID(ctx context.Context, taskID string) (*domain.Task, error)
	UpdateTaskByID(ctx context.Context, taskID string, updatedTask *domain.Task) error
	DeleteTaskByID(ctx context.Context, taskID string) error
}

type TaskRepositoryImpl struct {
	Collection *mongo.Collection
}
func NewTaskRepository(collection *mongo.Collection) TaskRepository {
	return &TaskRepositoryImpl{Collection: collection}
}
// function to create a new task in the database
func ( r *TaskRepositoryImpl) CreateTask(ctx context.Context,task *domain.Task) error{
	_,err:= r.Collection.InsertOne(ctx,task)
	return err
}

//function to get all tasks 
func (r *TaskRepositoryImpl) GetAllTasks(ctx context.Context) ([]domain.Task, error) {
    // Initialize empty slice to return empty slice in case of no tasks
    tasks := make([]domain.Task, 0)

    cur, err := r.Collection.Find(ctx, bson.M{})
    if err != nil {
        return nil, fmt.Errorf("failed to fetch tasks: %v", err)  
    }
    defer cur.Close(ctx)

    // Decode all 
    if err := cur.All(ctx, &tasks); err != nil {
        return nil, fmt.Errorf("failed to decode tasks: %v", err)
    }

    return tasks, nil
}
//function to get task by id
func (r *TaskRepositoryImpl) GetTaskByID(ctx context.Context, taskID string) (*domain.Task, error) {
	//check id 
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return nil, errors.New("invalid task ID")
	}

	//find task mapped with that id 

	var task domain.Task
	err = r.Collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, err
	}
	return &task, nil
}
//function to update task by id 

func (r *TaskRepositoryImpl) UpdateTaskByID(ctx context.Context, taskID string, updatedTask *domain.Task) error {
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return errors.New("invalid task ID")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"status":   updatedTask.Status,
		},
	}

	result, err := r.Collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return err
	}
	//count and if 0 it means no task
	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}
//function to delete task by id
func (r *TaskRepositoryImpl) DeleteTaskByID(ctx context.Context, taskID string) error {
	objID, err := primitive.ObjectIDFromHex(taskID)
	if err != nil {
		return errors.New("invalid task ID")
	}

	result, err := r.Collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}
	return nil
}