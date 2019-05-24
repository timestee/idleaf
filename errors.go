package idleaf

import (
	"errors"
)

var (
	errDomainLeafLost = errors.New("have no id leaf for the domain")
)

const (
	errOK         = 0
	errInternal   = 1
	errDomainLost = 2
)
