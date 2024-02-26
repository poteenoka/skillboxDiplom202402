package entity

import (
	"errors"
	"github.com/biter777/countries"
)

type MMSData struct {
	Country      string `json:"country"`
	Bandwith     string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func replaceCodeOnMMS(s *MMSData) {

	all := countries.AllInfo()
	for _, country := range all {
		if s.Country == country.Alpha2 {
			s.Country = country.Name
		}
	}

}

func (s *MMSData) ValidateMMSdata() error {
	Provider := [3]string{"Topolo", "Rond", "Kildy"}
	for _, item := range Provider {
		if item == s.Provider {
			return nil
		}
	}
	return errors.New("Not valid")
}

func (s *MMSData) SetCountry(counrty string) {
	s.Country = counrty
}
func (s *MMSData) GetCountry() string {
	return s.Country
}
func (s *MMSData) GetProvider() string {
	return s.Provider
}
