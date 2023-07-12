package post

type userImplementation struct {
	repo repository
}

func NewPostImplementation(repo repository) Service {
	return &userImplementation{
		repo: repo,
	}
}

type Service interface {
}
