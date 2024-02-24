package entity

import (
	"errors"
	"golang.org/x/text/language"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

func (s *EmailData) ValidateEmailData() error {
	_, err := language.Parse(s.Country)
	if err != nil {
		return err
	}
	var Provider = [13]string{"Gmail", "Gmail", "Hotmail", "MSN", "Orange", "Comcast", "AOL", "Live", "RediffMail", "GMX", "Protonmail",
		"Yandex", "Mail.ru"}

	for _, item := range Provider {
		if item == s.Provider {
			return nil
		}
	}
	return errors.New("Not valid")
}
