package repositories

import (
	"context"
	"linkupapi/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AvailabilityRepository struct {
	Collection *mongo.Collection
}

func NewAvailabilityRepository() *AvailabilityRepository {
	return &AvailabilityRepository{
		Collection: config.GetCollection("linkup", "availability"),
	}
}

func (r *AvailabilityRepository) CheckAvailability(ctx context.Context, zipCode string) (bool, error) {
	var result bson.M
	err := r.Collection.FindOne(ctx, bson.M{"zipcode": zipCode}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return false, nil
	}
	return err == nil, err
}
