package handler

type RegisterRequest struct {
	Fullname string `form:"fullname" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}
