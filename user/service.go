package user

type Service interface {
	FindAll() ([]User, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	return s.repository.FindAll()
}
