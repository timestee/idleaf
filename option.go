package idleaf

import (
	"github.com/timestee/goconf"
)

type Option struct {
	goconf.AutoOptions
	LeafTable     string `default:"id_leaf"`
	ServerPort    string `default:":16000"`
	IdOffset      int64  `default:"1000000"`
	DbProto       string `default:"mysql"`
	DbUser        string `default:"root"`
	DbPass        string `default:"123456"`
	DbHost        string `default:"127.0.0.1"`
	DbPort        int    `default:"3306"`
	DbName        string `default:"account"`
	TimeoutSecond int    `default:"5"`
}
