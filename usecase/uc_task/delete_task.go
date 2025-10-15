package uctask

func (tu *taskUsecase) DeleteTask(taskId uint) error {
	err := tu.taskRepository.DeleteTask(taskId)
	return err
}
