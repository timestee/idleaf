package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/timestee/goconf"
	"github.com/timestee/idleaf"
	"log"
	"net/http"
)

func main() {
	option := &idleaf.Option{}
	goconf.MustResolve(option)
	if err := idleaf.InitLeaf(option); err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(option.ServerPort, idleaf.InitRouter(option)))
}
