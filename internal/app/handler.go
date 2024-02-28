package app

import (
	"encoding/json"
	"github.com/skillboxDiplom202402/internal/entity"
	"net/http"
)

func HandleConnection(w http.ResponseWriter, r *http.Request) {

	structRes, err := GetResultData()
	if err != nil {
		answer := entity.ResultT{
			Status: false,
			Error:  "Не удалось получить данные. Сервис не доступен",
		}
		result, _ := json.Marshal(answer)
		w.Write(result)
		return
	}

	ans := entity.ResultT{
		Status: true,
		Data:   structRes,
	}

	resp, _ := json.Marshal(ans)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(resp)

}
