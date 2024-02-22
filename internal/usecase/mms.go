package usecase

type MmsService struct {
	repo csvReader
}

func NewMmsService(repo csvReader) *MmsService {
	return &MmsService{
		repo: repo,
	}
}

func (s *MmsService) SetData(data []string) error {
	return s.repo.SetData(data)
}

func (s *MmsService) GetContent(path string) ([]string, error) {
	return s.repo.GetContent(path)
}
