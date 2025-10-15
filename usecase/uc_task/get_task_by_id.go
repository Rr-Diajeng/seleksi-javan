package uctask

import "seleksi-javan/model/task"

func (tu *taskUsecase) GetTaskByID(taskId uint) (taskResp task.TaskResponse, err error) {
	t, err := tu.taskRepository.FindTaskByID(taskId)
	if err != nil {
		return taskResp, err
	}

	person, err := tu.userRepository.FindUserByID(t.AssignedID)
	if err != nil {
		return taskResp, err
	}

	taskResp = task.TaskResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		Username:    person.Username,
		Status:      string(t.Status),
	}

	return taskResp, nil
}
