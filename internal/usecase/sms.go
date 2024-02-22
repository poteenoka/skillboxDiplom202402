package usecase

import (
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"golang.org/x/text/language"
)

type SmsService struct {
	repo csvReader
}

func NewSmsService(repo csvReader) *SmsService {
	return &SmsService{
		repo: repo,
	}
}

func ValidateSMSVbs(s entity.SMSData) error {
	fmt.Println(s.Country)
	lang, err := language.Parse(s.Country)
	if err != nil {
		return err
	}
	fmt.Println(lang.String())
	return nil

}

func (s *SmsService) SetData(data []byte) error {
	return s.repo.SetData(data)
}

func (s *SmsService) GetContent(path byte) ([]string, error) {
	return s.repo.GetContent(path)
}
