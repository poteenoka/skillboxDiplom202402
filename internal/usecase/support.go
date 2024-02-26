package usecase

type SupportService struct {
	repo csvReader
}

func NewSupportService(repo csvReader) *SupportService {
	return &SupportService{
		repo: repo,
	}
}

func (s *SupportService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *SupportService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
