package entity

type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceData              `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   Billing                  `json:"billing"`
	Support   []SupportData            `json:"support"`
	Incidents []IncedentData           `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}
