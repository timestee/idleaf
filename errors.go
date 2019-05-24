package idleaf

import (
	"errors"
)

var (
	errDomainLeafLost = errors.New("have no id leaf for the domain")
)

const (
	ErrOK         = 0
	ErrInternal   = 1
	ErrDomainLost = 2
)
