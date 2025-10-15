package user

type UserResponse struct {
	ID       uint               `json:"id"`
	Username string             `json:"username"`
	Email    string             `json:"email"`
	Tasks    []TaskUserResponse `json:"tasks"`
}

type TaskUserResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
