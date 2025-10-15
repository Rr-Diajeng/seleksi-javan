package ucuser

import (
	"seleksi-javan/model"
	"seleksi-javan/model/user"
	"seleksi-javan/repository"
	"seleksi-javan/util/security"
)

type (
	UserUsecase interface {
		Register(userReq user.RegisterRequest) error
		Login(userReq user.LoginRequest) (*AuthResponsePayload, error)
		ChangePassword(userId uint, changePass user.ChangePasswordRequest) error
		GetAllUser() (userResp []user.UserResponse, err error)
		GetUserByID(userId uint) (userResp user.UserResponse, err error)
		UpdateUser(userId uint, updateUser user.UpdateUserRequest) error
		DeleteUser(userId uint) error
	}

	AuthResponsePayload struct {
		Token *security.GeneratedToken
		User  *model.User
	}

	userUsecase struct {
		userRepository repository.UserRepository
	}
)

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepo,
	}
}
