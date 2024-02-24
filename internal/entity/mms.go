package entity

import "errors"

type MMSData struct {
	Country      string `json:"country"`
	Bandwith     string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

/*func (s *MMSData) GetCountry() string {
	return s.Country
}
func (s *MMSData) GetProvider() string {
	return s.Provider
}
*/

func (s *MMSData) ValidateMMSdata() error {
	Provider := [3]string{"Topolo", "Rond", "Kildy"}
	for _, item := range Provider {
		if item == s.Provider {
			return nil
		}
	}
	return errors.New("Not valid")
}
