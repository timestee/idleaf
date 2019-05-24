package idleaf

import (
	"github.com/timestee/goconf"
)

// Option is the config for leaf
type Option struct {
	goconf.AutoOptions
	LeafTable     string `default:"id_leaf"`
	ServerPort    string `default:":16000"`
	IdOffset      int64  `default:"1000000"`
	DbProto       string `default:"mysql"`
	DbUser        string `default:"root"`
	DbPass        string `default:"111111"`
	DbHost        string `default:"127.0.0.1"`
	DbPort        int    `default:"3306"`
	DbName        string `default:"idleaf"`
	TimeoutSecond int    `default:"5"`
	BuffedCount   int64  `default:"2000"`
}
