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

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) createTeam(ctx context.Context, team Team) (Team, error) {
	args := m.Called(ctx, team)

	return args.Get(0).(Team), args.Error(1)
}
