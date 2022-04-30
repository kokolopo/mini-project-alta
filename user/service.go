package user

type UserService interface {
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) *userService {
	return &userService{repository}
}
