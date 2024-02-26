package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

/*func AppTestCollectData() {

	repoSms := repo.NewSMSLocalstorage()
	repoMms := repo.NewMMSLocalstorage()
	repoVoice := repo.NewVoiceLocalstorage()
	repoEmail := repo.NewEmailLocalstorage()
	repoBilling := repo.NewBillingLocalstorage()
	repoSupport := repo.NewSupportLocalstorage()
	repoIncedent := repo.NewIncedentLocalstorage()

	usecaseSms := usecase.NewSmsService(repoSms)
	var smsCSV []byte
	smsCSV, err := usecaseSms.GetContent(pathsmsDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseSms.SetData(smsCSV)
	repoSms.Print()

	usecaseMms := usecase.NewMmsService(repoMms)
	mmsCSV, err := usecaseMms.GetContent(pathmmsUrl)
	if err != nil {
		log.Fatal(err)
	}
	usecaseMms.SetData(mmsCSV)
	repoMms.Print()

	usecaseVoice := usecase.NewVoiceService(repoVoice)
	VoiceCSV, err := usecaseVoice.GetContent(pathvoiceDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseVoice.SetData(VoiceCSV)
	repoVoice.Print()

	usecaseEmail := usecase.NewVoiceService(repoEmail)
	emailCSV, err := usecaseEmail.GetContent(pathemailDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseEmail.SetData(emailCSV)
	repoEmail.Print()

	usecaseBilling := usecase.NewVoiceService(repoBilling)
	billingCSV, err := usecaseBilling.GetContent(pathbillingDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseBilling.SetData(billingCSV)
	//repoBilling.Print()

	usecaseSupport := usecase.NewSupportService(repoSupport)
	supportCSV, err := usecaseSupport.GetContent(pathsupportDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseSupport.SetData(supportCSV)
	repoSupport.Print()

	usecaseIncedent := usecase.NewSupportService(repoIncedent)
	incedentCSV, err := usecaseIncedent.GetContent(pathincedentDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseIncedent.SetData(incedentCSV)

}*/

func AppRun() {
	router := mux.NewRouter()

	GetResultData()

	router.HandleFunc("/", HandleConnection)
	srv := &http.Server{
		Addr:    "localhost:8282",
		Handler: router,
	}

	fmt.Println("Сервер запущен на порту 8282")
	srv.ListenAndServe()

}
