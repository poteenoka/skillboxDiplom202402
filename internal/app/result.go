package app

import (
	"github.com/biter777/countries"
	"github.com/skillboxDiplom202402/internal/entity"
	"github.com/skillboxDiplom202402/internal/usecase"
	"github.com/skillboxDiplom202402/internal/usecase/repo"
	"log"
	"time"
)

const pathsmsDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\sms.data"
const pathmmsUrl = "http://127.0.0.1:8383/mms"
const pathvoiceDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\voice.data"
const pathemailDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\email.data"
const pathbillingDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\billing.data"
const pathsupportDatafile = "http://127.0.0.1:8383/support"
const pathincedentDatafile = "http://127.0.0.1:8383/accendent"

var BufferedDataT entity.ResultSetT

func GetResultData() (t entity.ResultSetT, err error) {

	if BufferedDataT.Incidents != nil {
		return BufferedDataT, nil
	}

	repoSms := repo.NewSMSLocalstorage()
	repoMms := repo.NewMMSLocalstorage()
	repoVoice := repo.NewVoiceLocalstorage()
	repoEmail := repo.NewEmailLocalstorage()
	repoBilling := repo.NewBillingLocalstorage()
	repoSupport := repo.NewSupportLocalstorage()
	repoIncedent := repo.NewIncedentLocalstorage()

	usecaseSms := usecase.NewSmsService(repoSms)
	var smsCSV []byte
	smsCSV, err = usecaseSms.GetContent(pathsmsDatafile)
	if err != nil {
		return entity.ResultSetT{}, err
	}
	usecaseSms.SetData(smsCSV)

	usecaseMms := usecase.NewMmsService(repoMms)
	mmsCSV, err := usecaseMms.GetContent(pathmmsUrl)
	if err != nil {
		return entity.ResultSetT{}, err
	}
	usecaseMms.SetData(mmsCSV)

	repoMms.Print()

	usecaseVoice := usecase.NewVoiceService(repoVoice)
	VoiceCSV, err := usecaseVoice.GetContent(pathvoiceDatafile)
	if err != nil {
		return entity.ResultSetT{}, err
	}
	usecaseVoice.SetData(VoiceCSV)

	usecaseEmail := usecase.NewVoiceService(repoEmail)
	emailCSV, err := usecaseEmail.GetContent(pathemailDatafile)
	if err != nil {
		return entity.ResultSetT{}, err
	}
	usecaseEmail.SetData(emailCSV)
	//repoEmail.Print()

	usecaseBilling := usecase.NewVoiceService(repoBilling)
	billingCSV, err := usecaseBilling.GetContent(pathbillingDatafile)
	if err != nil {
		return entity.ResultSetT{}, err
	}
	usecaseBilling.SetData(billingCSV)

	usecaseSupport := usecase.NewSupportService(repoSupport)
	supportCSV, err := usecaseSupport.GetContent(pathsupportDatafile)
	if err != nil {
		log.Fatal(err)
		return entity.ResultSetT{}, err
	}
	usecaseSupport.SetData(supportCSV)

	usecaseIncedent := usecase.NewSupportService(repoIncedent)
	incedentCSV, err := usecaseIncedent.GetContent(pathincedentDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseIncedent.SetData(incedentCSV)

	sortedSms := prepareSMSdata(repoSms)
	//fmt.Println(sortedSms)

	chanDataMMS := make(chan [][]entity.MMSData, 1)
	prepareMMSdata(repoMms, chanDataMMS)
	dataMMS := <-chanDataMMS
	//fmt.Println("MMS DATA:...", dataMMS)

	chanDataEmail := make(chan map[string][][]entity.EmailData, 1)
	prepareEmailData(repoEmail, chanDataEmail)
	dataEmail := <-chanDataEmail

	//fmt.Println("Email Data:...", dataEmail)

	chanDataSupport := make(chan []int, 1)
	prepareSupportData(repoSupport, chanDataSupport)
	dataSupport := <-chanDataSupport
	//fmt.Println(dataSupport)

	chanDataIncident := make(chan []*entity.IncedentData, 1)
	prepareIncidentData(repoIncedent, chanDataIncident)
	dataIncidents := <-chanDataIncident
	//fmt.Println(dataIncidents[0])

	data := entity.ResultSetT{
		SMS:       sortedSms,
		MMS:       dataMMS,
		VoiceCall: repoVoice.Voice,
		Email:     dataEmail,
		Billing:   *repoBilling.Billing,
		Support:   dataSupport,
		Incidents: dataIncidents,
	}

	BufferedDataT = data
	return data, nil

}

func sortProvider(list []usecase.EntityData) []usecase.EntityData {
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list[j].GetProvider() > list[j+1].GetProvider() {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
	return list
}

func sortCountry(list []usecase.EntityData) []usecase.EntityData {
	// Bubble sort v.2 :)
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := length - 1; j > i; j-- {
			if list[j].GetCountry() < list[j-1].GetCountry() {
				list[j], list[j-1] = list[j-1], list[j]
			}
		}
	}
	return list
}

func prepareVoice(list []entity.VoiceData, ch chan []entity.VoiceData) {
	ch <- list
}

func prepareSMSdata(list *repo.SMSLocalstorage) [][]entity.SMSData {

	genMsg := make([]usecase.EntityData, 0, len(list.Sms))
	for i := range list.Sms {
		genMsg = append(genMsg, list.Sms[i])
	}

	sortedSMSByProvider := sortProvider(genMsg)
	sliceSMS1 := make([]entity.SMSData, len(sortedSMSByProvider))
	for i, message := range sortedSMSByProvider { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*entity.SMSData)
		if !ok {
		}
		sliceSMS1[i] = *smsMessage
		entity.ReplaceCodeOnSMS(&sliceSMS1[i])

	}

	sortedSMSByCountry := sortCountry(genMsg)
	sliceSMS2 := make([]entity.SMSData, len(sortedSMSByCountry))
	for i, message := range sortedSMSByCountry { // Приводим весь слайс []msgData к []SMSData
		smsMessage, ok := message.(*entity.SMSData)
		if !ok {
		}
		sliceSMS2[i] = *smsMessage
		entity.ReplaceCodeOnSMS(&sliceSMS2[i])
	}
	return [][]entity.SMSData{
		sliceSMS1, // second slice
		sliceSMS2, //first slice
	}
}

func prepareMMSdata(list *repo.MMSLocalstorage, ch chan [][]entity.MMSData) {
	genMsg := make([]usecase.EntityData, 0, len(list.Mms))
	for i := range list.Mms {
		genMsg = append(genMsg, list.Mms[i])
	}
	sortedMMSByProvider := sortProvider(genMsg)
	sliceMMS1 := make([]entity.MMSData, len(sortedMMSByProvider))
	for i, message := range sortedMMSByProvider {
		smsMessage, ok := message.(*entity.MMSData)
		if !ok {
		}
		sliceMMS1[i] = *smsMessage
	}
	sortedMMSByCountry := sortCountry(genMsg)
	sliceMMS2 := make([]entity.MMSData, len(sortedMMSByCountry))
	for i, message := range sortedMMSByCountry {
		smsMessage, ok := message.(*entity.MMSData)
		if !ok {
		}
		sliceMMS2[i] = *smsMessage
	}

	ch <- [][]entity.MMSData{
		sliceMMS1,
		sliceMMS2,
	}

}

func prepareEmailData(list *repo.EmailLocalstorage, ch chan map[string][][]entity.EmailData) {
	result := make(map[string][][]entity.EmailData)
	countryProvider := make(map[string][]entity.EmailData)

	//data.Country - поенять на имя

	all := countries.AllInfo()
	var cityname string

	for _, data := range list.Email {
		for _, country := range all {
			if data.Country == country.Alpha2 {
				cityname = country.Name
			}
		}
		countryProvider[cityname] = append(countryProvider[cityname], *data)
	}

	for country, provider := range countryProvider {
		sorted := sortMinMax(provider)
		result[country] = sorted
	}
	ch <- result
}

func sortMinMax(list []entity.EmailData) (result [][]entity.EmailData) {
	length := len(list)
	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list[j].DeliveryTime > list[j+1].DeliveryTime {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}

	//resultSort := sortDelivery(list.Email)

	min3 := make([]entity.EmailData, 3)
	max3 := make([]entity.EmailData, 3)
	min3 = list[:3]
	max3 = list[len(list)-3:]

	result = [][]entity.EmailData{
		min3,
		max3,
	}
	return result
}

func prepareSupportData(list *repo.SupportLocalstorage, ch chan []int) {
	result := make([]int, 2)
	hour := time.Now().Hour()
	if hour < 9 {
		result[0] = 1
	} else if hour > 16 {
		result[0] = 3
	} else {
		result[0] = 2
	}
	waitTime := func() int {
		result := 0
		for _, data1 := range list.Support {
			result += data1.ActiveTickets
		}
		//18tic/h
		return result * 18
	}()
	result[1] = waitTime
	ch <- result
}

func prepareIncidentData(list *repo.IncedentLocalstorage, ch chan []*entity.IncedentData) {

	length := len(list.Incedent)

	for i := 0; i < (length - 1); i++ {
		for j := 0; j < ((length - 1) - i); j++ {
			if list.Incedent[j+1].Status == "active" && list.Incedent[j].Status != "active" {
				list.Incedent[j], list.Incedent[j+1] = list.Incedent[j+1], list.Incedent[j]
			}
		}
	}

	ch <- list.Incedent
}
