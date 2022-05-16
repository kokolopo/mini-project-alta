package user

type InputRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type InputLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type InputCheckEmail struct {
	Email string `json:"email" binding:"required,email"`
}

type InputUpdate struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email" binding:"email"`
	Password string `json:"password"`
}
