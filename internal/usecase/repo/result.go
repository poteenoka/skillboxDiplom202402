package repo

import "github.com/skillboxDiplom202402/internal/entity"

type ResiultLocalstorage struct {
	ResultT *entity.ResultT
}

func NewResultLocalstorage() *ResiultLocalstorage {
	return &ResiultLocalstorage{
		ResultT: &entity.ResultT{
			Status: false,
			Data: entity.ResultSetT{
				SMS:       [][]entity.SMSData{},
				MMS:       [][]entity.MMSData{},
				VoiceCall: []entity.VoiceData{},
				Email:     map[string][][]entity.EmailData{},
				Billing:   entity.Billing{},
				Support:   []entity.SupportData{},
				Incidents: []entity.IncedentData{},
			},
			Error: "",
		},
	}
}
