package main

import (
	"github.com/timestee/goconf"
	"github.com/timestee/idleaf"
)

func main() {
	ops := &idleaf.Option{}
	goconf.MustResolve(ops)

	idleaf.Init(ops)
	router := idleaf.InitRouter()
	router.Run(ops.ServerPort)
}
