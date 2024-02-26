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

type SupportLocalstorage struct {
	Support []*entity.SupportData
}

func NewSupportLocalstorage() *SupportLocalstorage {
	return &SupportLocalstorage{
		Support: make([]*entity.SupportData, 0),
	}
}

func (s *SupportLocalstorage) Print() {
	for _, data := range s.Support {
		fmt.Printf("Ticket: %d\n", data.ActiveTickets)
		fmt.Printf("topik: %s\n", data.Topic)
	}
}

func (s *SupportLocalstorage) GetContent(path string) ([]byte, error) {

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

	return nil, errors.New("что-то пошло не так")

}

func (s *SupportLocalstorage) SetData(body []byte) error {

	//var SupportBuf []entity.SupportData

	err := json.Unmarshal(body, &s.Support)
	if err != nil {
		log.Println(err)
		return err
	}

	//for _, v := range SupportBuf {
	//	if v.ValidateSupportdata() == nil {
	//		s.Support = append(s.Support, &v)
	//	}
	//}

	return nil
}
