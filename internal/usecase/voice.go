package usecase

type VoiceService struct {
	repo csvReader
}

func NewVoiceService(repo csvReader) *VoiceService {
	return &VoiceService{
		repo: repo,
	}
}

func (s *VoiceService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *VoiceService) GetContent(path string) ([]byte, error) {
	return s.repo.GetContent(path)
}
