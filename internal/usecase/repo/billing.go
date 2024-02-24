package repo

import (
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type BillingLocalstorage struct {
	Billing *entity.Billing
}

func NewBillingLocalstorage() *BillingLocalstorage {
	return &BillingLocalstorage{
		Billing: &entity.Billing{
			CreateCustomer: false,
			Purchase:       false,
			Payout:         false,
			Recurring:      false,
			FraudControl:   false,
			CheckoutPage:   false,
		},
	}
}

func (s *BillingLocalstorage) Print() {
	fmt.Println(s.Billing)
}

func (s *BillingLocalstorage) GetContent(path string) ([]byte, error) {
	// Считываем CSV-файл в []byte
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	return data, nil
}

func (s *BillingLocalstorage) SetData(data []byte) error {
	stringData := string(data)
	var buffer float64
	for i := 0; i < len(stringData); i++ {
		index := len(stringData) - 1 - i
		bit := string(stringData[index])
		if bit == "1" {
			buffer = buffer + math.Pow(2, float64(i))
		}
	}
	decData := uint8(buffer)
	binData := strconv.FormatInt(int64(decData), 2)

	length := len(binData)
	binData = strings.Repeat("0", 6-length) + binData
	boolMask := make([]bool, 6)
	for i := 0; i < 6; i++ {
		if binData[i:i+1] == "1" {
			boolMask[i] = true
		}
	}
	s.Billing.CreateCustomer = boolMask[5]
	s.Billing.Purchase = boolMask[4]
	s.Billing.Payout = boolMask[3]
	s.Billing.Recurring = boolMask[2]
	s.Billing.FraudControl = boolMask[1]
	s.Billing.CheckoutPage = boolMask[0]

	return nil
}
