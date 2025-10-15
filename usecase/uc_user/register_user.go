package ucuser

import (
	"net/http"
	"seleksi-javan/model"
	"seleksi-javan/model/user"
	"seleksi-javan/util/http/errors"
	"seleksi-javan/util/security"

	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) Register(userReq user.RegisterRequest) error {
	isUsernameExist, err := uu.userRepository.FindUserByUsername(userReq.Username)
	if err == nil && isUsernameExist.ID != 0 {
		return errors.NewHttpError(http.StatusConflict, "Username already exists")
	}

	isEmailExist, err := uu.userRepository.FindUserByEmail(userReq.Email)
	if err == nil && isEmailExist.ID != 0 {
		return errors.NewHttpError(http.StatusConflict, "Email already exists")
	}

	hashedPassword, err := security.HashPassword(userReq.Password)
	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return errors.NewHttpError(
				http.StatusBadRequest,
				"Password is too long. Please use fewer characters.",
			)
		}

		return err
	}

	newUser := model.User{
		Username: userReq.Username,
		Email:    userReq.Email,
		Password: hashedPassword,
	}

	err = uu.userRepository.CreateUser(newUser)
	if err != nil {
		return err
	}

	return nil

}
