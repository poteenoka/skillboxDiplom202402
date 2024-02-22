package main

import (
	"github.com/skillboxDiplom202402/internal/usecase"
	"github.com/skillboxDiplom202402/internal/usecase/repo"
	"log"
)

func main() {

	const pathsmsDatafile = "d:\\github\\skillboxDiplom202402\\simulator\\sms.data"
	const pathsmsUrl = "http://127.0.0.1:8383/mms"

	repoSms := repo.NewSMSLocalstorage()
	usecaseSms := usecase.NewSmsService(repoSms)
	var smsCSV []string
	smsCSV, err := usecaseSms.GetContent(pathsmsDatafile)
	if err != nil {
		log.Fatal(err)
	}
	usecaseSms.SetData(smsCSV)
	//repoSms.Print()

	repoMms := repo.NewMMSLocalstorage()
	becauseMms := usecase.NewMmsService(repoMms)

	becauseMms.GetContent(pathsmsUrl)
	repoMms.Print()

	//smsSt := NewUserService
	//var smsFromFile []string
	//smsSt.
	//smsFromFile, err := usecase.GetFileContent(path.Join(pathDatafile, "sms.data"))
	//
	//if err != nil {
	//	log.Fatalln("файла с данными не существует")
	//
	//}
	//fmt.Println(smsFromFile)

}
