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
		return errors.New("model.Movie.ValidateFields(): movie name can't be empty")
	}

	if strings.TrimSpace(m.Director) == "" {
		return errors.New("model.Movie.ValidateFields(): movie director can't be empty")
	}

	if len(m.MainCast) == 0 {
		return errors.New("model.Movie.ValidateFields(): movie main cast can't be empty")
	}

	for _, c := range m.MainCast {
		if strings.TrimSpace(c) == "" {
			return errors.New("model.Movie.ValidateFields(): main cast name can't be empty")
		}
	}

	return nil
}
