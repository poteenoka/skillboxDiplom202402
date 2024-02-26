package entity

import (
	"github.com/biter777/countries"
	"golang.org/x/text/language"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwith     string `json:"bandwith"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func ReplaceCodeOnSMS(s *SMSData) {

	all := countries.AllInfo()
	for _, country := range all {
		if s.Country == country.Alpha2 {
			s.Country = country.Name
		}
	}

}

func ValidateSMSVbs(s SMSData) error {
	_, err := language.Parse(s.Country)
	if err != nil {
		return err
	}
	//fmt.Println(lang.String())
	return nil
}

func (s *SMSData) SetCountry(counrty string) {
	s.Country = counrty
}
func (s *SMSData) GetCountry() string {
	return s.Country
}
func (s *SMSData) GetProvider() string {
	return s.Provider
}
