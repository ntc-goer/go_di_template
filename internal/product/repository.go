package product

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_di_template/pkg/database"
)

type RepositoryI interface {
	GetAll(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]*Product, error)
}
type Repository struct {
	Collection database.CollectionInterface
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Collection: db.Collection(PRODUCT_COLLECTION_NAME),
	}
}
func (r Repository) GetAll(ctx context.Context, filter bson.M, opts *options.FindOptions) ([]*Product, error) {
	//TODO implement me
	cur, err := r.Collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	var results []*Product
	if err = cur.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return results, nil
}
