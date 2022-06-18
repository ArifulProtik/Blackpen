package services

import (
	"github.com/ArifulProtik/BlackPen/ent"
)

type Service struct {
	Auth *AuthService
	User *UserService
}

func New(dbclient *ent.Client) *Service {
	return &Service{
		Auth: &AuthService{
			Client: dbclient.Auth,
		},
		User: &UserService{
			Client: dbclient.User,
		},
	}
}
