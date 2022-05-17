package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(input InputRegister) (User, error)
	Login(input InputLogin) (User, error)
	IsEmailAvailable(input InputCheckEmail) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	UpdateUser(id int, input InputUpdate) (User, error)
	GetUserById(id int) (User, error)
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
	newUser.Fullname = input.Fullname
	newUser.Email = input.Email
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

	//cek jika user tidak ada
	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *userService) IsEmailAvailable(input InputCheckEmail) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (s *userService) UpdateUser(id int, input InputUpdate) (User, error) {
	user, errUser := s.repository.FindById(id)
	if errUser != nil {
		return user, errUser
	}

	user.ID = id
	user.Fullname = input.Fullname
	user.Email = input.Email

	//enkripsi password
	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if errHash != nil {
		return user, errHash
	}
	user.Password = string(passwordHash)

	updatedUser, errUpdate := s.repository.UpdateByID(user)
	if errUpdate != nil {
		return updatedUser, errUpdate
	}

	return updatedUser, nil
}

func (s *userService) SaveAvatar(id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	user.Avatar = fileLocation

	updatedUser, errUpdate := s.repository.Update(user)
	if errUpdate != nil {
		return updatedUser, errUpdate
	}

	return updatedUser, nil
}

func (s *userService) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("no user found")
	}

	return user, nil
}
