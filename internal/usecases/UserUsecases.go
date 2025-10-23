package usecases

import (
	"github.com/google/uuid"
	"github.com/stacoviaki/api-mave/internal/models"
	"github.com/stacoviaki/api-mave/internal/repositories"
)

// 🔹 UserUseCase
// Representa a camada de regras de negócio (intermediária entre Controller e Repository).
// Aqui ficam as lógicas que fazem sentido para o domínio da aplicação.
type UserUseCase struct {
	repositories repositories.UserRepositories // acesso ao banco via repositório
}

// 🔹 NewUserUseCase
// Construtor da camada de UseCase.
// Recebe um repositório já conectado e devolve o UseCase pronto pra uso.
func NewUserUseCase(repo repositories.UserRepositories) UserUseCase {
	return UserUseCase{
		repositories: repo,
	}
}

// 🔹 GetUsers
// Retorna todos os usuários cadastrados.
// Apenas repassa a chamada para o repositório, sem regras extras.
func (us *UserUseCase) GetUsers() ([]models.User, error) {
	return us.repositories.GetUsers()
}

// 🔹 GetUserById
// Busca um usuário específico pelo seu UUID.
// Se o usuário não existir, retorna nil (sem erro).
func (us *UserUseCase) GetUserById(id_user uuid.UUID) (*models.User, error) {
	user, err := us.repositories.GetUserById(id_user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// 🔹 UpdateUser
// Atualiza os dados de um usuário existente.
// Recebe o UUID e o novo objeto de usuário, e retorna o usuário atualizado.
func (us *UserUseCase) UpdateUser(id_user uuid.UUID, updatedUser models.User) (*models.User, error) {
	// Chama a função de UPDATE no repositório
	user, err := us.repositories.UpdateUser(id_user, updatedUser)
	if err != nil {
		return nil, err
	}

	// Retorna o usuário já atualizado
	return user, nil
}

// 🔹 CreateUser
// Cria um novo usuário no banco de dados.
// Primeiro insere no repositório, depois adiciona o ID retornado ao struct.
func (us *UserUseCase) CreateUser(user models.User) (models.User, error) {
	// Cria o usuário no banco e recebe o UUID gerado automaticamente
	userId, err := us.repositories.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	// Adiciona o ID ao struct e retorna
	user.ID = userId
	return user, nil
}

// 🔹 DeleteUser
// Exclui um usuário do banco com base no UUID.
// Retorna o ID deletado em caso de sucesso.
func (us *UserUseCase) DeleteUser(id_user uuid.UUID) (uuid.UUID, error) {
	deletedID, err := us.repositories.DeleteUser(id_user)
	if err != nil {
		return uuid.Nil, err
	}
	return deletedID, nil
}
