package mongodb

import (
	"context"
	"time"

	"github.com/erkylima/hexagonal/internal/beneficiary"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	client   *mongo.Client
	database string
	timeout  time.Duration
}

func newMongoClient(connectionString string, timeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}
	// err = client.Ping(ctx, readpref.Primary())
	// if err != nil {
	// 	return nil, err
	// }

	return client, err
}

func NewMongoRepository(mongoUrl, mongoDB string, mongoTimeout int) (beneficiary.BeneficiaryRepository, error) {
	repo := &mongoRepository{
		timeout:  time.Duration(mongoTimeout) * time.Second,
		database: mongoDB,
	}

	client, err := newMongoClient(mongoUrl, mongoTimeout)

	if err != nil {
		return nil, err
	}

	repo.client = client
	return repo, nil
}

func (db *mongoRepository) Find(name string) (*beneficiary.Beneficiary, error) {
	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()
	benefic := &beneficiary.Beneficiary{}
	collection := db.client.Database(db.database).Collection("beneficiary")
	filter := bson.M{"name": name}
	err := collection.FindOne(ctx, filter).Decode(&benefic)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.Wrap(beneficiary.ErrBeneficiaryNotFound, "repository.Beneficiary.Find")
		}
		return nil, errors.Wrap(err, "repository.Beneficiary.Find")
	}
	return benefic, nil
}

func (db *mongoRepository) Store(beneficiary *beneficiary.Beneficiary) error {
	ctx, cancel := context.WithTimeout(context.Background(), db.timeout)
	defer cancel()
	collection := db.client.Database(db.database).Collection("beneficiary")
	_, err := collection.InsertOne(ctx, bson.M{
		"name":    beneficiary.Name,
		"age":     beneficiary.Age,
		"address": beneficiary.Address,
		"phone":   beneficiary.Phone,
	})
	if err != nil {
		return errors.Wrap(err, "repository.Beneficiary.Store")
	}
	return nil
}
