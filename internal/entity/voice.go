package entity

import (
	"errors"
	"golang.org/x/text/language"
)

type VoiceData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

func (s *VoiceData) ValidateVoiceDatadata() error {
	_, err := language.Parse(s.Country)
	if err != nil {
		return err
	}
	var Provider = [3]string{"TransparentCalls", "E-Voice", "JustPhone"}
	for _, item := range Provider {
		if item == s.Provider {
			return nil
		}
	}
	return errors.New("Not valid")
}
