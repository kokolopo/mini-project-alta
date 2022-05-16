package user

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

var repoUser = &RepositoryMock{Mock: mock.Mock{}}
var serviceUser = userService{repository: repoUser}

func TestGetUserById(t *testing.T) {
	user := User{
		ID:        1,
		Fullname:  "fahmi",
		Email:     "fahmi@gmail.com",
		Password:  "password",
		Role:      "admin",
		Avatar:    "avatar.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoUser.Mock.On("FindById", 12).Return(user)

	result, err := serviceUser.GetUserById(12)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Fullname, result.Fullname)
}

// func TestRegister(t *testing.T) {
// 	var input = InputRegister{
// 		Fullname: "fahmi",
// 		Email:    "fahmi@gmail.com",
// 		Password: "password",
// 	}
// 	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
// 	user := User{
// 		Fullname:  "fahmi",
// 		Email:     "fahmi@gmail.com",
// 		Password:  string(passwordHash),
// 		Role:      "user",
// 		Avatar:    "",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	repoUser.Mock.On("Save", user).Return(user)

// 	result, err := serviceUser.Register(input)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// 	assert.Equal(t, user.ID, result.ID)
// 	assert.Equal(t, user.Fullname, result.Fullname)
// }

func TestLogin(t *testing.T) {
	var input = InputLogin{Email: "fahmi@gmail.com", Password: "password"}
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user := User{
		ID:        1,
		Fullname:  "fahmi",
		Email:     "fahmi@gmail.com",
		Password:  string(passwordHash),
		Role:      "admin",
		Avatar:    "avatar.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoUser.Mock.On("FindByEmail", "fahmi@gmail.com").Return(user)

	result, err := serviceUser.Login(input)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Fullname, result.Fullname)
}

func TestIsEmailAvailable(t *testing.T) {
	var input = InputCheckEmail{Email: "fahmi@gmail.com"}
	user := User{
		ID:        1,
		Fullname:  "fahmi",
		Email:     "fahmi@gmail.com",
		Password:  "password",
		Role:      "admin",
		Avatar:    "avatar.png",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoUser.Mock.On("FindByEmail", input.Email).Return(user)

	result, err := serviceUser.IsEmailAvailable(input)
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

func TestSaveAvatar(t *testing.T) {
	user := User{
		ID:        1,
		Fullname:  "fahmi",
		Email:     "fahmi@gmail.com",
		Password:  "password",
		Role:      "admin",
		Avatar:    "images/avatar.jpg",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoUser.Mock.On("FindById", 1).Return(user)
	repoUser.Mock.On("Update", user).Return(user)

	result, err := serviceUser.SaveAvatar(1, "images/avatar.jpg")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}

// func TestUpdateUser(t *testing.T) {
// 	var input = InputUpdate{
// 		Fullname: "fahmi",
// 		Email:    "fahmi@gmail.com",
// 		Password: "password",
// 	}
// 	user := User{
// 		ID:        1,
// 		Fullname:  "fahmi",
// 		Email:     "fahmi@gmail.com",
// 		Password:  "password",
// 		Role:      "admin",
// 		Avatar:    "images/avatar.jpg",
// 		CreatedAt: time.Now(),
// 		UpdatedAt: time.Now(),
// 	}

// 	repoUser.Mock.On("FindById", 1).Return(user)
// 	repoUser.Mock.On("UpdateByID", user).Return(user)

// 	result, err := serviceUser.UpdateUser(1, input)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// }
func TestUpdateUserFail(t *testing.T) {
	input := InputUpdate{}
	user := User{}

	repoUser.Mock.On("FindById", 0).Return(nil)
	repoUser.Mock.On("UpdateByID", user).Return(nil)

	_, err := serviceUser.UpdateUser(0, input)

	assert.Nil(t, nil)
	assert.NotNil(t, err)
}
