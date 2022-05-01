package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input InputRegister) (User, error)
	Login(input InputLogin) (User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *userService {
	return &userService{repository}
}

func (s *userService) Register(input InputRegister) (User, error) {
	var newUser User

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return newUser, errHash
	}

	//tangkap nilai dari inputan
	newUser.NamaLengkap = input.NamaLengkap
	newUser.Email = input.Email
	newUser.Whatsapp = input.Whatsapp
	newUser.Password = string(passwordHash)
	newUser.Role = "user"

	//save data yang sudah dimapping kedalam struct Mahasiswa
	user, err := s.repository.Save(newUser)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) Login(input InputLogin) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}
