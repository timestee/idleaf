package idleaf

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/timestee/goconf"
)

var router *mux.Router
var testMysql = false

func TestMain(m *testing.M) {
	ops := &Option{}
	goconf.MustResolve(ops)
	err := Init(ops)
	if err == nil {
		testMysql = true
	}
	router = InitRouter(ops)
	m.Run()
}

func gen(t testing.TB, domain string) {
	if !testMysql {
		return
	}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/gen/"+domain, nil)
	router.ServeHTTP(w, req)
	idRet := &resp{}
	err := json.NewDecoder(w.Body).Decode(idRet)
	if err != nil {
		t.Fail()
	}
	if idRet.Code != 0 {
		t.Fail()
	}
}

func TestGenId(t *testing.T) {
	gen(t, "test")
}

func BenchmarkGenIdDomainOid(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gen(b, "oid")
		}
	})
}

func BenchmarkGenIdDomainUid(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			gen(b, "uid")
		}
	})
}
