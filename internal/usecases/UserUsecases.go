package usecases

import (
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/repositories"
)

type UserUseCase struct {
	repositories repositories.UserRepositories
}

func NewUserUseCase(repo repositories.UserRepositories) UserUseCase {
	return UserUseCase{
		repositories: repo,
	}
}

func (us *UserUseCase) GetUsers() ([]models.User, error) {
	return us.repositories.GetUsers()
}

func (us *UserUseCase) GetUserById(id_user uuid.UUID) (*models.User, error) {
	user, err := us.repositories.GetUserById(id_user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserUseCase) CreateUser(user models.User) (models.User, error) {
	userId, err := us.repositories.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}
	user.ID = userId
	return user, nil
}
