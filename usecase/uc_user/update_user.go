package ucuser

import (
	"net/http"
	"seleksi-javan/model"
	"seleksi-javan/model/user"
	"seleksi-javan/util/http/errors"
	"seleksi-javan/util/security"

	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) UpdateUser(userId uint, userReq user.UpdateUserRequest) error {
	updateData := model.User{}

	if userReq.Username != nil {
		updateData.Username = *userReq.Username
	}

	if userReq.Email != nil {
		updateData.Email = *userReq.Email
	}

	if userReq.Password != nil {
		password, err := security.HashPassword(*userReq.Password)

		if err != nil {
			if err == bcrypt.ErrPasswordTooLong {
				return errors.NewHttpError(
					http.StatusBadRequest,
					"Password is too long. Please use fewer characters.",
				)
			}

			return err
		}

		updateData.Password = password
	}

	err := uu.userRepository.UpdateUser(userId, updateData)
	return err
}
