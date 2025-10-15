package uctask

import (
	"seleksi-javan/model"
	"seleksi-javan/model/task"
)

func (tu *taskUsecase) AddTask(taskReq task.TaskRequest) error {
	newTask := model.Task{
		Title:       taskReq.Title,
		Description: taskReq.Description,
		Status:      model.StatusTask(taskReq.Status),
		AssignedID:  taskReq.AssignedID,
	}

	err := tu.taskRepository.CreateTask(newTask)
	return err
}
