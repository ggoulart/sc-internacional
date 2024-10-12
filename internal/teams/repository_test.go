package teams

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestRepository_createTeam(t *testing.T) {
	objectId := primitive.NewObjectID()
	tests := []struct {
		name    string
		setup   func(d *dbMock)
		team    Team
		want    Team
		wantErr error
	}{
		{
			name: "when failed to create a team",
			setup: func(d *dbMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				d.On("InsertOne", mock.Anything, receivedTeam, []*options.InsertOneOptions(nil)).Return(&mongo.InsertOneResult{}, errors.New("failed to create team"))
			},
			team:    Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			want:    Team{},
			wantErr: errors.New("failed to create team"),
		},
		{
			name: "when successfully create a team",
			setup: func(d *dbMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				d.On("InsertOne", mock.Anything, receivedTeam, []*options.InsertOneOptions(nil)).Return(&mongo.InsertOneResult{InsertedID: objectId}, nil)
			},
			team:    Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			want:    Team{Id: objectId.Hex(), Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbMock{}
			tt.setup(d)

			r := NewRepository(d)

			got, err := r.createTeam(context.Background(), tt.team)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRepository_getTeam(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(d *dbMock)
		id      string
		want    Team
		wantErr error
	}{
		{
			name:    "when invalid id is received",
			setup:   func(d *dbMock) {},
			id:      "xpto",
			want:    Team{},
			wantErr: errors.New("the provided hex string is not a valid ObjectID"),
		},
		{
			name: "when failed to find team",
			setup: func(d *dbMock) {
				hexId := primitive.ObjectID{0x67, 0xa, 0x95, 0xa8, 0xc1, 0x35, 0xef, 0x7c, 0x3d, 0x61, 0xf3, 0xb5}
				d.On("FindOne", mock.Anything, primitive.M{"_id": hexId}, []*options.FindOneOptions(nil)).Return(mongo.NewSingleResultFromDocument(nil, nil, nil))
			},
			id:      "670a95a8c135ef7c3d61f3b5",
			want:    Team{},
			wantErr: errors.New("document is nil"),
		},
		{
			name: "when sucessfully find team",
			setup: func(d *dbMock) {
				hexId := primitive.ObjectID{0x67, 0xa, 0x95, 0xa8, 0xc1, 0x35, 0xef, 0x7c, 0x3d, 0x61, 0xf3, 0xb5}
				result := map[string]interface{}{"_id": "670a95a8c135ef7c3d61f3b5", "name": "Internacional", "fullName": "Sport Club Internacional", "website": "internacional.com.br", "foundationDate": time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				d.On("FindOne", mock.Anything, primitive.M{"_id": hexId}, []*options.FindOneOptions(nil)).Return(mongo.NewSingleResultFromDocument(result, nil, nil))
			},
			id:      "670a95a8c135ef7c3d61f3b5",
			want:    Team{Id: "670a95a8c135ef7c3d61f3b5", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbMock{}
			tt.setup(d)

			r := NewRepository(d)

			got, err := r.getTeam(context.Background(), tt.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestRepository_getAllTeams(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(d *dbMock)
		want    []Team
		wantErr error
	}{
		{
			name: "when failed to find teams",
			setup: func(d *dbMock) {
				d.On("Find", mock.Anything, bson.M{}, []*options.FindOptions(nil)).Return(&mongo.Cursor{}, errors.New("failed to find"))
			},
			want:    []Team{},
			wantErr: errors.New("failed to find"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dbMock{}
			tt.setup(d)

			r := NewRepository(d)

			got, err := r.getAllTeams(context.Background())

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

type dbMock struct {
	db
	mock.Mock
}

func (m *dbMock) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, document, opts)

	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *dbMock) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	args := m.Called(ctx, filter, opts)

	return args.Get(0).(*mongo.SingleResult)
}

func (m *dbMock) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (cur *mongo.Cursor, err error) {
	args := m.Called(ctx, filter, opts)

	return args.Get(0).(*mongo.Cursor), args.Error(1)
}
