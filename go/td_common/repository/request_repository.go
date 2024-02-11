package repository

import (
	"context"

	"github.com/Tracking-Detector/td-backend/go/td_common/config"
	"github.com/Tracking-Detector/td-backend/go/td_common/model"
	"github.com/Tracking-Detector/td-backend/go/td_common/mongostore"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRequestRepository struct {
	db   *mongo.Database
	coll *mongo.Collection
}

func NewMongoRequestRepository(db *mongo.Database) *MongoRequestRepository {
	coll := db.Collection(config.EnvRequestCollection())
	mongostore.EnsureIndex(context.Background(), coll, "dataset", 1)
	return &MongoRequestRepository{
		db:   db,
		coll: coll,
	}
}

func (r *MongoRequestRepository) Save(ctx context.Context, m *model.RequestData) (*model.RequestData, error) {
	return mongostore.Save(ctx, r.coll, m)
}

func (r *MongoRequestRepository) SaveAll(ctx context.Context, m []*model.RequestData) ([]*model.RequestData, error) {
	return mongostore.SaveAll(ctx, r.coll, m)
}

func (r *MongoRequestRepository) FindAll(ctx context.Context) ([]*model.RequestData, error) {
	return mongostore.FindAll(ctx, r.coll, &model.RequestData{})
}

func (r *MongoRequestRepository) FindAllByUrlLikePaged(ctx context.Context, url string, page, pageSize int) ([]*model.RequestData, error) {
	findOptions := options.Find()
	filter := bson.M{}
	if url != "" {
		filter = bson.M{
			"url": bson.M{
				"$regex": primitive.Regex{
					Pattern: url,
					Options: "i",
				},
			},
		}
	}
	findOptions.SetSkip((int64(page) - 1) * int64(pageSize))
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64((page - 1) * pageSize))
	findOptions.SetLimit(int64(pageSize))

	requestDataValues, err := mongostore.FindAllBy(ctx, r.coll, &model.RequestData{}, filter, findOptions)
	if err != nil {
		return nil, err
	}

	return requestDataValues, nil
}

func (r *MongoRequestRepository) StreamAll(ctx context.Context) (<-chan *model.RequestData, <-chan error) {
	return mongostore.StreamAll[*model.RequestData](ctx, r.coll, bson.M{})
}

func (r *MongoRequestRepository) StreamByDataset(ctx context.Context, dataset string) (<-chan *model.RequestData, <-chan error) {
	return mongostore.StreamAll[*model.RequestData](ctx, r.coll, bson.M{
		"dataset": dataset,
	})
}

func (r *MongoRequestRepository) FindByID(ctx context.Context, id string) (*model.RequestData, error) {
	return mongostore.FindByID(ctx, r.coll, id, &model.RequestData{})
}

func (r *MongoRequestRepository) Count(ctx context.Context) (int64, error) {
	return mongostore.Count(ctx, r.coll)
}

func (r *MongoRequestRepository) CountByUrlLike(ctx context.Context, url string) (int64, error) {
	filter := bson.M{}
	if url != "" {
		filter = bson.M{
			"url": bson.M{
				"$regex": primitive.Regex{
					Pattern: url,
					Options: "i",
				},
			},
		}
	}
	return mongostore.CountBy(ctx, r.coll, filter)
}

func (r *MongoRequestRepository) DeleteByID(ctx context.Context, id string) error {
	return mongostore.DeleteByID(ctx, r.coll, id)
}

func (r *MongoRequestRepository) DeleteAll(ctx context.Context) error {
	return mongostore.DeleteAll(ctx, r.coll)
}

func (r *MongoRequestRepository) DeleteAllByLabel(ctx context.Context, label string) error {
	return mongostore.DeleteAllBy(ctx, r.coll, bson.M{
		"dataset": label,
	})
}

func (r *MongoRequestRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	return mongostore.InTransaction(ctx, r.db, fn)
}
