package ucuser

import (
	"net/http"
	"seleksi-javan/model/user"
	"seleksi-javan/util/http/errors"
	"seleksi-javan/util/security"

	"golang.org/x/crypto/bcrypt"
)

func (uu *userUsecase) ChangePassword(userId uint, changePass user.ChangePasswordRequest) error {
	user, err := uu.userRepository.FindUserByID(userId)
	if err != nil {
		return err
	}

	if !security.CheckPasswordHash(changePass.OldPassword, user.Password) {
		return errors.NewHttpError(
			http.StatusUnauthorized,
			"Password does not match",
		)
	}

	newPassword, err := security.HashPassword(changePass.NewPassword)

	if err != nil {
		if err == bcrypt.ErrPasswordTooLong {
			return errors.NewHttpError(
				http.StatusBadRequest,
				"Password is too long. Please use fewer characters.",
			)
		}

		return err
	}

	return uu.userRepository.ChangePassword(userId, newPassword)
}
