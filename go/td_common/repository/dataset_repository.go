package repository

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/mongostore"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoDatasetRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewMongoDatasetRepository(db *mongo.Database) *MongoDatasetRepository {
	coll := db.Collection(config.EnvDatasetCollection())
	return &MongoDatasetRepository{
		db:   db,
		coll: coll,
	}
}

func (r *MongoDatasetRepository) Save(ctx context.Context, m *model.Dataset) (*model.Dataset, error) {
	return mongostore.Save(ctx, r.coll, m)
}

func (r *MongoDatasetRepository) SaveAll(ctx context.Context, m []*model.Dataset) ([]*model.Dataset, error) {
	return mongostore.SaveAll(ctx, r.coll, m)
}

func (r *MongoDatasetRepository) StreamAll(ctx context.Context) (<-chan *model.Dataset, <-chan error) {
	return mongostore.StreamAll[*model.Dataset](ctx, r.coll, bson.M{})
}

func (r *MongoDatasetRepository) FindByID(ctx context.Context, id string) (*model.Dataset, error) {
	return mongostore.FindByID(ctx, r.coll, id, &model.Dataset{})
}

func (r *MongoDatasetRepository) FindAll(ctx context.Context) ([]*model.Dataset, error) {
	return mongostore.FindAll(ctx, r.coll, &model.Dataset{})
}

func (r *MongoDatasetRepository) FindByName(ctx context.Context, name string) (*model.Dataset, error) {
	return mongostore.FindByName(ctx, r.coll, name, &model.Dataset{})
}

func (r *MongoDatasetRepository) FindByLabel(ctx context.Context, label string) (*model.Dataset, error) {
	return mongostore.FindBy(ctx, r.coll, bson.M{
		"label": label,
	}, &model.Dataset{})
}

func (r *MongoDatasetRepository) DeleteAll(ctx context.Context) error {
	return mongostore.DeleteAll(ctx, r.coll)
}

func (r *MongoDatasetRepository) DeleteByID(ctx context.Context, id string) error {
	return mongostore.DeleteByID(ctx, r.coll, id)
}

func (r *MongoDatasetRepository) Count(ctx context.Context) (int64, error) {
	return mongostore.Count(ctx, r.coll)
}

func (r *MongoDatasetRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return mongostore.InTransaction(ctx, r.db, fn)
}
