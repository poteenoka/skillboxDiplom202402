package repo

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/skillboxDiplom202402/internal/entity"
	"io/ioutil"
	"strconv"
	"strings"
)

type VoiceLocalstorage struct {
	Voice []*entity.VoiceData
}

func (s *VoiceLocalstorage) Print() {
	for _, VoiceData := range s.Voice {
		fmt.Println(VoiceData)
	}
}

func NewVoiceLocalstorage() *VoiceLocalstorage {
	return &VoiceLocalstorage{
		Voice: make([]*entity.VoiceData, 0),
	}
}

func (s *VoiceLocalstorage) GetContent(path string) ([]byte, error) {
	// Считываем CSV-файл в []byte
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return nil, err
	}
	return data, nil
}

func (s *VoiceLocalstorage) SetData(data []byte) error {

	buffer := bytes.NewBuffer(data)
	scanner := bufio.NewScanner(buffer)
	for scanner.Scan() {
		spliArr := strings.Split(scanner.Text(), ";")
		//слайс должен быть длинной 8
		if len(spliArr) != 8 {
			continue
		}

		ConnectionStability, err := strconv.ParseFloat(spliArr[4], 32)
		if err != nil {
			continue
		}
		TTFB, err := strconv.Atoi(spliArr[5])
		if err != nil {
			continue
		}
		VoicePurity, err := strconv.Atoi(spliArr[6])
		if err != nil {
			continue
		}
		MedianOfCallsTime, err := strconv.Atoi(spliArr[7])
		if err != nil {
			continue
		}

		Voice := entity.VoiceData{
			Country:             spliArr[0],
			Bandwidth:           spliArr[1],
			ResponseTime:        spliArr[2],
			Provider:            spliArr[3],
			ConnectionStability: float32(ConnectionStability),
			TTFB:                TTFB,
			VoicePurity:         VoicePurity,
			MedianOfCallsTime:   MedianOfCallsTime}

		err = Voice.ValidateVoiceDatadata()
		if err != nil {
			continue
		}

		s.Voice = append(s.Voice, &Voice)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения строк:", err)
		return err
	}
	return nil
}
