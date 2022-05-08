package user

type UserFormatter struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
}

func FormatUser(user User, token string) UserFormatter {
	formatter := UserFormatter{
		Id:       user.ID,
		Name:     user.Fullname,
		Email:    user.Email,
		Password: user.Password,
		Avatar:   user.Avatar,
		Token:    token,
	}

	return formatter
}
