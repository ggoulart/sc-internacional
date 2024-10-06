package teams

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestService_CreateTeam(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(r *repositoryMock)
		team    Team
		want    Team
		wantErr error
	}{
		{
			name: "when repository fail to create team",
			setup: func(r *repositoryMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				r.On("createTeam", mock.Anything, receivedTeam).Return(Team{}, errors.New("failed to create team"))
			},
			team:    Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			want:    Team{},
			wantErr: errors.New("failed to create team"),
		},
		{
			name: "when repository successfully create team",
			setup: func(r *repositoryMock) {
				receivedTeam := Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				returnTeam := Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				r.On("createTeam", mock.Anything, receivedTeam).Return(returnTeam, nil)
			},
			team:    Team{Id: "", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			want:    Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repositoryMock{}
			tt.setup(r)

			s := NewService(r)

			got, err := s.createTeam(context.Background(), tt.team)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func TestService_GetTeamById(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(r *repositoryMock)
		id      string
		want    Team
		wantErr error
	}{
		{
			name: "when repository fail to get team",
			setup: func(r *repositoryMock) {
				r.On("getTeam", mock.Anything, "1").Return(Team{}, errors.New("failed to get team"))
			},
			id:      "1",
			want:    Team{},
			wantErr: errors.New("failed to get team"),
		},
		{
			name: "when repository successfully get team",
			setup: func(r *repositoryMock) {
				returnTeam := Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)}
				r.On("getTeam", mock.Anything, "1").Return(returnTeam, nil)
			},
			id:      "1",
			want:    Team{Id: "1", Name: "Internacional", FullName: "Sport Club Internacional", Website: "internacional.com.br", FoundationDate: time.Date(1909, time.April, 4, 0, 0, 0, 0, time.UTC)},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repositoryMock{}
			tt.setup(r)

			s := NewService(r)

			got, err := s.getTeam(context.Background(), tt.id)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

type repositoryMock struct {
	repository
	mock.Mock
}

func (m *repositoryMock) createTeam(ctx context.Context, team Team) (Team, error) {
	args := m.Called(ctx, team)

	return args.Get(0).(Team), args.Error(1)
}

func (m *repositoryMock) getTeam(ctx context.Context, id string) (Team, error) {
	args := m.Called(ctx, id)

	return args.Get(0).(Team), args.Error(1)
}
