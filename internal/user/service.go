package user

import (
	"context"
	"rest-api-server/pkg/logging"
)

type service struct {
	storage Storage
	logger  *logging.Logger
}

func (s *service) Create(ctx context.Context, dto CreateUserDTO) (u User, err error) {
	return
}
