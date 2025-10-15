package task

type (
	TaskRequest struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
		AssignedID  uint   `json:"assigned_id" binding:"required"`
		Status      string `json:"status" binding:"required,oneof=pending in_progress completed"`
	}

	TaskUpdateRequest struct {
		Title       *string `json:"title" binding:"required"`
		Description *string `json:"description" binding:"required"`
		AssignedID  *uint   `json:"assigned_id" binding:"required"`
		Status      *string `json:"status" binding:"required,oneof=pending in_progress completed"`
	}
)
