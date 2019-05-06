package idleaf

import (
	"errors"
)

var (
	ErrDomainLeafLost = errors.New("have no id leaf for the domain")
)

const (
	ErrOK       = 0
	ErrInternal = 1
)
