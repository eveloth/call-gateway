package models

import (
	"errors"
	"strings"
)

type email struct {
	Address string
}

func NewEmail(emailAddress string) (*email, error) {
	if strings.TrimSpace(emailAddress) == "" {
		return nil, errors.New("Email should not be empty")
	}

	if !strings.Contains(emailAddress, "@") {
		return nil, errors.New("Not a valid email address")
	}

	return &email{Address: emailAddress}, nil
}
