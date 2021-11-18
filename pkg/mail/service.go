package mail

type repository interface {
	Send(em Email) (Email, error)
}

type Service struct {
	repo repository
}

func NewService(repo repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Send(e Email) (Email, error) {
	return s.repo.Send(e)
}
