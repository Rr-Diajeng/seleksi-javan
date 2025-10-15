package uctask

import (
	"seleksi-javan/model/task"
	"seleksi-javan/repository"
)

type (
	TaskUsecase interface {
		AddTask(taskReq task.TaskRequest) error
		GetAllTask() (taskResp []task.TaskResponse, err error)
		GetTaskByID(taskId uint) (taskResp task.TaskResponse, err error)
		UpdateTask(taskId uint, updateTask task.TaskUpdateRequest) error
		DeleteTask(taskId uint) error
	}

	taskUsecase struct {
		taskRepository repository.TaskRepository
		userRepository repository.UserRepository
	}
)

func NewTaskUsecase(taskRepo repository.TaskRepository, userRepo repository.UserRepository) TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepo,
		userRepository: userRepo,
	}
}
