package idleaf

import (
	"encoding/json"
	"github.com/timestee/goconf"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var router *gin.Engine

func TestMain(m *testing.M) {
	ops := &Option{}
	goconf.MustResolve(ops)
	Init(ops)
	router = InitRouter()
	m.Run()
}

func genid(t testing.TB, domain string) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/gen/"+domain, nil)
	router.ServeHTTP(w, req)
	idRet := &struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Id   int64  `json:"id"`
	}{}
	err := json.NewDecoder(w.Body).Decode(idRet)
	if err != nil {
		t.Fail()
	}
	if idRet.Code != 0 {
		t.Fail()
	}
}

func TestGenId(t *testing.T) {
	genid(t, "test")
}

func BenchmarkGenIdDomainOid(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			genid(b, "oid")
		}
	})
}

func BenchmarkGenIdDomainUid(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			genid(b, "uid")
		}
	})
}
