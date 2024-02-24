package repo

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"io/ioutil"
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

func (s *SMSLocalstorage) GetContent(path string) ([]byte, error) {

	// Считываем CSV-файл в []byte
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}

	return data, nil
}

func (s *SMSLocalstorage) SetData(data []byte) error {

	buffer := bytes.NewBuffer(data)

	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		spliArr := strings.Split(scanner.Text(), ";")
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

		if entity.ValidateSMSVbs(sms) != nil {
			continue
		}
		s.Sms = append(s.Sms, &sms)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения строк:", err)
		return err
	}

	return nil
}
