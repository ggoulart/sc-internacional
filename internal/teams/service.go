package teams

import "context"

type repository interface {
	createTeam(ctx context.Context, team Team) (Team, error)
	getTeam(ctx context.Context, id string) (Team, error)
	getAllTeams(ctx context.Context) ([]Team, error)
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

func (s Service) getTeam(ctx context.Context, id string) (Team, error) {
	team, err := s.repository.getTeam(ctx, id)
	if err != nil {
		return Team{}, err
	}

	return team, nil
}

func (s Service) getAllTeams(ctx context.Context) ([]Team, error) {
	teams, err := s.repository.getAllTeams(ctx)
	if err != nil {
		return nil, err
	}

	return teams, nil
}
