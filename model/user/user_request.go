package user

type (
	RegisterRequest struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,max=72"`
	}

	LoginRequest struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required,min=8,max=72"`
	}

	ChangePasswordRequest struct {
		OldPassword string `json:"old_password" binding:"required,min=8,max=72"`
		NewPassword string `json:"new_password" binding:"required,min=8,max=72"`
	}

	UpdateUserRequest struct {
		Username *string `json:"username,omitempty"`
		Email    *string `json:"email,omitempty"`
		Password *string `json:"password,omitempty"`
	}
)
