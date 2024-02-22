package repo

import (
	"bufio"
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"github.com/skillboxDiplom202402/internal/usecase"
	"log"
	"os"
	"strings"
)

type SMSLocalstorage struct {
	Sms []*entity.SMSData
}

func (s *SMSLocalstorage) Print() {
	for _, smsData := range s.Sms {
		fmt.Printf("Страна: %s\n", smsData.Country)
		fmt.Printf("Пропускная способность: %s\n", smsData.Bandwith)
		fmt.Printf("Время отклика: %s\n", smsData.ResponseTime)
		fmt.Printf("Провайдер: %s\n\n", smsData.Provider)
	}
}

func NewSMSLocalstorage() *SMSLocalstorage {
	return &SMSLocalstorage{
		Sms: make([]*entity.SMSData, 0),
	}
}

func (s *SMSLocalstorage) GetContent(path string) ([]string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("file does not exist")
			return nil, err
		}
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		rows = append(rows, sc.Text())
	}
	return rows, nil

}

func (s *SMSLocalstorage) SetData(data []string) error {

	for _, row := range data {
		spliArr := strings.Split(row, ";")

		//слайс должен быть длинной 4
		if len(spliArr) != 4 {
			continue
		}
		sms := entity.SMSData{
			Country:      spliArr[0],
			Bandwith:     spliArr[1],
			ResponseTime: spliArr[2],
			Provider:     spliArr[3],
		}

		if usecase.ValidateSMSVbs(sms) != nil {
			continue
		}

		s.Sms = append(s.Sms, &sms)
	}
	return nil
}
