package api

import "github.com/dsbasko/yandex-go-diploma-1/services/auth/internal/entities"

type ChangePasswordRequestV1 struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangePasswordResponseV1 entities.RepositoryAccountEntity

/*
Трансформация структур
*/

func ChangePasswordResponseV1FromEntity(entity *entities.RepositoryAccountEntity) *ChangePasswordResponseV1 {
	return &ChangePasswordResponseV1{
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
