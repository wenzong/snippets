package user

import (
	"context"

	"github.com/gocraft/dbr/v2"
	"github.com/wenzong/demo/api/pb"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (u *Service) Get(ctx context.Context, userID int64) (*pb.User, error) {
	return u.repo.Get(userID)
}

func (u *Service) TxGet(ctx context.Context, tx *dbr.Tx, userID int64) (*pb.User, error) {
	return u.repo.TxGet(tx, userID)
}
