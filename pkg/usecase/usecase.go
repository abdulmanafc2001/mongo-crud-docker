package usecase

import (
	"crud/pkg/models"
	"crud/pkg/repos"
	"log"
)

type UserUsecase interface {
	CreateUser(models.User) (any, error)
	ReadUser() (any, error)
	DeleteUsers() error
	GetOneUser(string) (any, error)
	UpdateUser(string, models.User) error
	DeleteUser(string) error
}

type UserUseCase struct {
	UseCase repos.UserRepositoryInterface
}

func NewUsecase(usecase repos.UserRepositoryInterface) UserUsecase {
	return &UserUseCase{
		UseCase: usecase,
	}
}

func (u *UserUseCase) CreateUser(usr models.User) (any, error) {
	id, err := u.UseCase.CreateUser(usr)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return id, nil
}

func (u *UserUseCase) ReadUser() (any, error) {
	usr, err := u.UseCase.ReadUser()

	if err != nil {
		return nil, err
	}

	return usr, nil

}

func (u *UserUseCase) DeleteUsers() error {
	return u.UseCase.DeleteUsers()
}

func (u *UserUseCase) GetOneUser(id string) (any, error) {
	return u.UseCase.GetOneUser(id)
}

func (u *UserUseCase) UpdateUser(id string, usr models.User) error {
	return u.UseCase.UpdateUser(id, usr)
}
func (u *UserUseCase) DeleteUser(id string) error {
	return u.UseCase.DeleteUser(id)
}
