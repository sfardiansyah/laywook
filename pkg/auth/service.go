package auth

import "errors"

var (
	// ErrInvalidCredentials ...
	ErrInvalidCredentials = errors.New("Invalid username or password")
)

// Service ...
type Service interface {
	GetUser(string) (*User, error)
}

// Repository ...
type Repository interface {
	GetUser(string) (*User, error)
}

type service struct {
	r Repository
}

// NewService creates an auth service with the necessary dependencies.
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetUser(email string) (*User, error) {
	user, err := s.r.GetUser(email)
	if err != nil {
		return nil, err
	}

	return user, err
}
