package main

import (
	"github.com/timestee/goconf"
	"github.com/timestee/idleaf"
	"log"
	"net/http"
)

func main() {
	option := &idleaf.Option{}
	goconf.MustResolve(option)
	idleaf.MustCheckError(idleaf.Init(option))
	log.Fatal(http.ListenAndServe(option.ServerPort, idleaf.InitRouter()))
}
