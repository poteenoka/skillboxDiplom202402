package repo

import (
	"errors"
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"io"
	"log"
	"net/http"
)

type MMSLocalstorage struct {
	Mms []*entity.MMSData
}

func NewMMSLocalstorage() *MMSLocalstorage {
	return &MMSLocalstorage{
		Mms: make([]*entity.MMSData, 0),
	}
}

func (s *MMSLocalstorage) Print() {
	for _, mmsData := range s.Mms {
		fmt.Printf("Страна: %s\n", mmsData.Country)
		fmt.Printf("Пропускная способность: %s\n", mmsData.Bandwith)
		fmt.Printf("Время отклика: %s\n", mmsData.ResponseTime)
		fmt.Printf("Провайдер: %s\n\n", mmsData.Provider)
	}
}

func (s *MMSLocalstorage) GetContent(path string) ([]byte, error) {

	resp, err := http.Get(path)

	if err != nil {
		log.Println(err)
		log.Println("ощибка  ", path)
	}
	if resp.StatusCode != 200 {
		log.Println("Status code is not 200, error is occured")
		return nil, errors.New("Status code is not 200, error is occured")
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return body, nil
	}

	//for i, v := range s.Mms {
	//	if v.Country {
	//		list = append(list[:i], list[i+1:]...)
	//	}
	//}

	return nil, nil

}

func (s *MMSLocalstorage) SetData([]byte) error {
	return nil
}
