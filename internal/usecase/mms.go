package usecase

type MmsService struct {
	repo csvReader
}

func NewMmsService(repo csvReader) *MmsService {
	return &MmsService{
		repo: repo,
	}
}

func (s *MmsService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *MmsService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
