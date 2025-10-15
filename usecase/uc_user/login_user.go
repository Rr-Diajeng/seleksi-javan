package ucuser

import (
	"net/http"
	"seleksi-javan/model/user"
	"seleksi-javan/util/http/errors"
	"seleksi-javan/util/security"
)

func (uu *userUsecase) Login(userReq user.LoginRequest) (*AuthResponsePayload, error) {
	user, err := uu.userRepository.FindUserByUsername(userReq.Username)
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.NewHttpError(
			http.StatusUnauthorized,
			"No data found",
		)
	}

	if !security.CheckPasswordHash(userReq.Password, user.Password) {
		return nil, errors.NewHttpError(
			http.StatusUnauthorized,
			"Email or password does not match",
		)
	}

	generatedToken, err := security.GenerateToken(&user)
	if err != nil {
		return nil, err
	}

	return &AuthResponsePayload{
		Token: generatedToken,
		User:  &user,
	}, nil
}
