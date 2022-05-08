package user

type InputRegister struct {
	Fullname string `json:"nama_lengkap" binding:"required"`
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
	ID       int
	Fullname string `json:"nama_lengkap"`
	Email    string `json:"email" binding:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Error    error
}
