package usecase

type EmailService struct {
	repo csvReader
}

func NewEmailService(repo csvReader) *EmailService {
	return &EmailService{
		repo: repo,
	}
}

func (s *EmailService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *EmailService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
