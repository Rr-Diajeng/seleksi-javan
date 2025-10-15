package uctask

import (
	"fmt"
	"seleksi-javan/model"
	"seleksi-javan/model/task"
)

func (tu *taskUsecase) UpdateTask(taskId uint, updateTask task.TaskUpdateRequest) error {
	existingTask := model.Task{}

	if updateTask.Title != nil {
		existingTask.Title = *updateTask.Title
	}

	if updateTask.Description != nil {
		existingTask.Description = *updateTask.Description
	}

	if updateTask.AssignedID != nil {
		existingTask.AssignedID = *updateTask.AssignedID
	}

	if updateTask.Status != nil {
		status := model.StatusTask(*updateTask.Status)
		if status != model.Pending && status != model.InProgress && status != model.Completed {
			return fmt.Errorf("invalid status value")
		}
		existingTask.Status = status
	}

	if err := tu.taskRepository.UpdateTask(taskId, existingTask); err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	return nil
}
