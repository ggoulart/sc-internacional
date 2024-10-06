package teams

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(database *mongo.Database) *Repository {
	return &Repository{collection: database.Collection("teams")}
}

func (r Repository) createTeam(ctx context.Context, team Team) (Team, error) {
	result, err := r.collection.InsertOne(ctx, team)
	if err != nil {
		return Team{}, err
	}

	team.Id = result.InsertedID.(primitive.ObjectID).Hex()

	return team, nil
}

func (r Repository) getTeam(ctx context.Context, id string) (Team, error) {
	var team Team
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Team{}, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": docID}).Decode(&team)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}
