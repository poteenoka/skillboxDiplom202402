package usecase

type IncedentService struct {
	repo csvReader
}

func NewIncedentService(repo csvReader) *IncedentService {
	return &IncedentService{
		repo: repo,
	}
}

func (s *IncedentService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *IncedentService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
