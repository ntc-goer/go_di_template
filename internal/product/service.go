package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceI interface {
	GetAll(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]*Product, error)
}

type Service struct {
	Repository RepositoryI
}

func NewService(r *Repository) *Service {
	return &Service{
		Repository: r,
	}
}

func (s *Service) GetAll(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]*Product, error) {
	//TODO implement me
	return s.Repository.GetAll(ctx, filter, opts)
}
