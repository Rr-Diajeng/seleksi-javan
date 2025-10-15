package ucuser

import "seleksi-javan/model/user"

func (uu *userUsecase) GetAllUser() (userResp []user.UserResponse, err error) {
	users, err := uu.userRepository.GetAllUser()
	if err != nil {
		return nil, err
	}

	for _, u := range users {

		var taskResponses []user.TaskUserResponse
		for _, t := range u.Tasks {
			taskResponses = append(taskResponses, user.TaskUserResponse{
				ID:          t.ID,
				Title:       t.Title,
				Description: t.Description,
				Status:      string(t.Status),
			})
		}

		userResp = append(userResp, user.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
			Tasks:    taskResponses,
		})
	}

	return userResp, nil
}
