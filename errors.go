package idleaf

import (
	"errors"
	"log"
)

var (
	ErrDomainLeafLost = errors.New("have no id leaf for the domain")
)

const (
	ErrOK         = 0
	ErrInternal   = 1
	ErrDomainLost = 2
)

func MustCheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
