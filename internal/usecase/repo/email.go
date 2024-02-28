package repo

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"io/ioutil"
	"strconv"
	"strings"
)

type EmailLocalstorage struct {
	Email []*entity.EmailData
}

func (s *EmailLocalstorage) Print() {
	for _, EmailData := range s.Email {
		fmt.Println(EmailData)
		fmt.Println("-------", EmailData.Country)
	}
}

func NewEmailLocalstorage() *EmailLocalstorage {
	return &EmailLocalstorage{
		Email: make([]*entity.EmailData, 0),
	}
}

func (s *EmailLocalstorage) GetContent(path string) ([]byte, error) {
	// Считываем CSV-файл в []byte
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	return data, nil
}

func (s *EmailLocalstorage) SetData(data []byte) error {

	buffer := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		spliArr := strings.Split(scanner.Text(), ";")
		//слайс должен быть длинной 3
		if len(spliArr) != 3 {
			continue
		}
		DeliveryTime, err := strconv.Atoi(spliArr[2])
		if err != nil {
			continue
		}
		Email := entity.EmailData{
			Country:      spliArr[0],
			Provider:     spliArr[1],
			DeliveryTime: DeliveryTime}

		err = Email.ValidateEmailData()
		if err != nil {
			continue
		}
		s.Email = append(s.Email, &Email)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения строк:", err)
		return err
	}
	return nil
}
