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

type IncedentLocalstorage struct {
	Incedent []*entity.IncedentData
}

func NewIncedentLocalstorage() *IncedentLocalstorage {
	return &IncedentLocalstorage{
		Incedent: make([]*entity.IncedentData, 0),
	}
}

func (s *IncedentLocalstorage) Print() {
	for _, data := range s.Incedent {
		fmt.Printf("status: %s\n", data.Status)
		fmt.Printf("topik: %s\n", data.Topic)
	}
}

func (s *IncedentLocalstorage) GetContent(path string) ([]byte, error) {

	resp, err := http.Get(path)
	if err != nil {
		log.Println(err)
		log.Println("ошибка  ", path)
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

func (s *IncedentLocalstorage) SetData(body []byte) error {

	//var IncedentBuf []entity.IncedentData

	err := json.Unmarshal(body, &s.Incedent)
	if err != nil {
		log.Println(err)
		return err
	}

	//for _, v := range IncedentBuf {
	//	if v.ValidateIncedentdata() == nil {
	//		s.Incedent = append(s.Incedent, &v)
	//	}
	//}

	return nil
}
