package teams

import "context"

type repository interface {
	createTeam(ctx context.Context, team Team) (Team, error)
}

type Service struct {
	repository repository
}

func NewService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s Service) createTeam(ctx context.Context, team Team) (Team, error) {
	createdTeam, err := s.repository.createTeam(ctx, team)
	if err != nil {
		return Team{}, err
	}

	return createdTeam, nil
}
