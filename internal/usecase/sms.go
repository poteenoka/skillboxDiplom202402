package usecase

type SmsService struct {
	repo csvReader
}

func NewSmsService(repo csvReader) *SmsService {
	return &SmsService{
		repo: repo,
	}
}

func (s *SmsService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *SmsService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
