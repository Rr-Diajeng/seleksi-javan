package task

type TaskResponse struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Username    string `json:"assigned_to"`
	Status      string `json:"status"`
}
