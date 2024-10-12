package teams

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type db interface {
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error)
}

type Repository struct {
	db
}

func NewRepository(db db) *Repository {
	return &Repository{db}
}

func (r Repository) createTeam(ctx context.Context, team Team) (Team, error) {
	result, err := r.db.InsertOne(ctx, team)
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

	err = r.db.FindOne(ctx, bson.M{"_id": docID}).Decode(&team)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}

func (r Repository) getAllTeams(ctx context.Context) ([]Team, error) {
	cursor, err := r.db.Find(ctx, bson.M{})
	if err != nil {
		return []Team{}, err
	}
	defer cursor.Close(ctx)

	var teams []Team

	err = cursor.All(ctx, &teams)
	if err != nil {
		return []Team{}, err
	}

	return teams, nil
}
