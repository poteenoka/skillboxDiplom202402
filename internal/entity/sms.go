package entity

import (
	"golang.org/x/text/language"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwith     string `json:"bandwith"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

/*
func (s *SMSData) GetCountry() string {
	return s.Country
}
func (s *SMSData) GetProvider() string {
	return s.Provider
}*/

func ValidateSMSVbs(s SMSData) error {
	_, err := language.Parse(s.Country)
	if err != nil {
		return err
	}
	//fmt.Println(lang.String())
	return nil
}
