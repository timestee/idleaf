package idleaf

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
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

func InitRouter(option *Option) *mux.Router {
	BuffedCount = option.BuffedCount
	withTimeout := func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), time.Duration(option.TimeoutSecond)*time.Second)
			defer cancel()
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/v1/gen/{domain}", withTimeout(http.HandlerFunc(GenDomainId)))
	return router
}
