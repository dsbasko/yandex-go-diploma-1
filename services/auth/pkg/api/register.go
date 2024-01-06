package api

import "github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/entities"

type RegisterRequestV1 struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegisterResponseV1 entities.RepositoryAccountEntity

/*
Трансформация структур
*/

func RegisterResponseV1FromEntity(entity *entities.RepositoryAccountEntity) *RegisterResponseV1 {
	return &RegisterResponseV1{
		ID:        entity.ID,
		Username:  entity.Username,
		Password:  entity.Password,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		LastLogin: entity.LastLogin,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}
