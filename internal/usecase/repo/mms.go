package repo

import (
	"encoding/json"
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
		log.Printf("Error http req: %d", resp.StatusCode)
		return nil, errors.New(resp.Status)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err == nil {
		return body, nil
	}

	//for i, v := range s.Mms {
	//	if v.Country {
	//		list = append(list[:i], list[i+1:]...)
	//	}
	//}

	return nil, errors.New("что-то пошло не так")

}

func (s *MMSLocalstorage) SetData(body []byte) error {
	var mmsBuf []entity.MMSData

	err := json.Unmarshal(body, &mmsBuf)
	if err != nil {
		log.Println(err)
		return err
	}

	for _, v := range mmsBuf {
		if v.ValidateMMSdata() == nil {
			s.Mms = append(s.Mms, &v)
		}
	}
	return nil
}
