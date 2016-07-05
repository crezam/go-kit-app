package businesslogic

import (
	"errors"
	"strings"
)

type StringService interface {
	UpperCase(string) (string, error)
	Count(string) int
}

type StringServiceImpl struct {
}

func (StringServiceImpl) UpperCase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (StringServiceImpl) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")
