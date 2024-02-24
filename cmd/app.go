package main

import (
	"github.com/skillboxDiplom202402/internal/usecase"
	"github.com/skillboxDiplom202402/internal/usecase/repo"
	"log"
)

func appRun() {
	const pathsmsDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\sms.data"
	const pathmmsUrl = "http://127.0.0.1:8383/mms"
	const pathvoiceDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\voice.data"
	const pathemailDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\email.data"
	const pathbillingDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\billing.data"

	repoSms := repo.NewSMSLocalstorage()
	repoMms := repo.NewMMSLocalstorage()
	repoVoice := repo.NewVoiceLocalstorage()
	repoEmail := repo.NewEmailLocalstorage()
	repoBilling := repo.NewBillingLocalstorage()

	usecaseSms := usecase.NewSmsService(repoSms)
	var smsCSV []byte

	smsCSV, err := usecaseSms.GetContent(pathsmsDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseSms.SetData(smsCSV)
	//repoSms.Print()
	usecaseMms := usecase.NewMmsService(repoMms)
	mmsCSV, err := usecaseMms.GetContent(pathmmsUrl)
	if err != nil {
		log.Fatal(err)
	}
	usecaseMms.SetData(mmsCSV)
	//repoMms.Print()

	usecaseVoice := usecase.NewVoiceService(repoVoice)
	VoiceCSV, err := usecaseVoice.GetContent(pathvoiceDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseVoice.SetData(VoiceCSV)
	//repoVoice.Print()

	usecaseEmail := usecase.NewVoiceService(repoEmail)
	emailCSV, err := usecaseEmail.GetContent(pathemailDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseEmail.SetData(emailCSV)
	// repoEmail.Print()

	usecaseBilling := usecase.NewVoiceService(repoBilling)
	billingCSV, err := usecaseBilling.GetContent(pathbillingDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseBilling.SetData(billingCSV)
	repoBilling.Print()

}
