package entity

type MMSData struct {
	Country      string `json:"country"`
	Bandwith     string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

func (s *MMSData) GetCountry() string {
	return s.Country
}
func (s *MMSData) GetProvider() string {
	return s.Provider
}
