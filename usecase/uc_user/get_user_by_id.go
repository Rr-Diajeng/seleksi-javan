package ucuser

import "seleksi-javan/model/user"

func (uu *userUsecase) GetUserByID(userId uint) (userResp user.UserResponse, err error) {
	usr, err := uu.userRepository.FindUserByID(userId)
	if err != nil {
		return userResp, err
	}
	if usr.ID == 0 {
		return userResp, nil
	}

	var taskResponses []user.TaskUserResponse
	for _, t := range usr.Tasks {
		taskResponses = append(taskResponses, user.TaskUserResponse{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      string(t.Status),
		})
	}

	userResp = user.UserResponse{
		ID:       usr.ID,
		Username: usr.Username,
		Email:    usr.Email,
		Tasks:    taskResponses,
	}

	return userResp, nil
}
