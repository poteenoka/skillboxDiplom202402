package usecase

type BillingService struct {
	repo csvReader
}

func NewBillingService(repo csvReader) *BillingService {
	return &BillingService{
		repo: repo,
	}
}

func (s *BillingService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *BillingService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
