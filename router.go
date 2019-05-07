package idleaf

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Id   int64  `json:"id"`
}

func jsonResp(resp *Resp, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	payload, err := json.Marshal(resp)
	if err != nil {
		payload = []byte(`{"code":1,"msg":"json marshal error","id":0}`)
	}
	_, _ = w.Write(payload)
}

func GenDomainId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	domain, ok := vars["domain"]
	rsp := &Resp{Code: ErrInternal}
	if !ok || domain == "" {
		rsp.Code = ErrDomainLost
		rsp.Msg = "domain lost"
	} else {
		if id, err := idLeaf.GenId(domain); err == nil {
			rsp.Code = ErrOK
			rsp.Id = id
		} else {
			rsp.Code = ErrInternal
			rsp.Msg = err.Error()
		}
	}
	jsonResp(rsp, w)
}

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/gen/{domain}", GenDomainId)
	return router
}
