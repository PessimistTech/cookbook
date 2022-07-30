package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recipe struct {
	ID           ObjectID          `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string            `json:"title"`
	Metadata     map[string]string `json:"metadata"`
	Ingregients  []*Ingredient     `json:"ingredients"`
	Instructions []*Instruction    `json:"instructions"`
	Notes        []*string         `json:"notes,omitempty"`
}

type Ingredient struct {
	Quantity   int    `json:"quantity"`
	Unit       string `json:"unit"`
	Ingredient string `json:"ingredient"`
}

type Instruction struct {
	StepNum     int    `json:"stepNum"`
	Instruction string `json:"instruction"`
}

const (
	database   = "cookbook"
	collection = "recipes"
)

func (r *Recipe) Delete() error {
	coll := GetCollection(database, collection)
	objectId, err := primitive.ObjectIDFromHex(string(r.ID))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	res, err := coll.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	if res.DeletedCount < 1 {
		return fmt.Errorf("no object deleted")
	}

	return nil
}
