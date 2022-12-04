package bo

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Identity int    `json:"identity"`
	Token    string `json:"token"`
}

type ApplyRegisterRequest struct {
	Email string `json:"email" binding:"required"`
}

type ApplyRegisterResponse struct{}

type ActiveUserRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
	Password   string `json:"password" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Department string `json:"department" binding:"required"`
}

type ActiveUserResponse struct{}

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}

type ResetPasswordResponse struct{}

type ActiveResetPasswordRequest struct {
	Email      string `json:"email" binding:"required"`
	ActiveCode string `json:"active_code" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type ActiveResetPasswordResponse struct{}
