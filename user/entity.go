package user

import "time"

type User struct {
	ID          int
	NamaLengkap string
	Email       string
	Whatsapp    string
	Password    string
	Role        string
	Avatar      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
