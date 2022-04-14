package model

import (
	"errors"
	"strings"
)

type Movie struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director string   `json:"director"`
	MainCast []string `json:"main_cast"`
}

func (m *Movie) ValidateFields() error {
	if strings.TrimSpace(m.Name) == "" {
		return errors.New("movie name can't be empty")
	}

	if strings.TrimSpace(m.Director) == "" {
		return errors.New("movie director can't be empty")
	}

	if len(m.MainCast) == 0 {
		return errors.New("movie main cast can't be empty")
	}

	for _, c := range m.MainCast {
		if strings.TrimSpace(c) == "" {
			return errors.New("main cast name can't be empty")
		}
	}

	return nil
}
