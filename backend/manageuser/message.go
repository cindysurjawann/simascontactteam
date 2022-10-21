package manageuser

type GetUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type UpdateUserRequest struct {
	Email string `json:"email" binding:"required"`
	Role  int8   `json:"role" binding:"required"`
	Name  string `json:"name" binding:"required"`
}