package uctask

import "seleksi-javan/model/task"

func (tu *taskUsecase) GetAllTask() (taskResp []task.TaskResponse, err error) {
	tasks, err := tu.taskRepository.GetAllTask()
	if err != nil {
		return nil, err
	}

	for _, t := range tasks {
		person, err := tu.userRepository.FindUserByID(t.AssignedID)
		if err != nil {
			return nil, err
		}

		taskResp = append(taskResp, task.TaskResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Username:    person.Username,
			Status:      string(t.Status),
		})
	}

	return taskResp, nil
}
