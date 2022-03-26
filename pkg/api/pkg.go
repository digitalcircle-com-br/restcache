package api

import (
	"net/http"
	"time"

	"github.com/digitalcircle-com-br/restdb/pkg/infra/db"
	"github.com/digitalcircle-com-br/restdb/pkg/infra/router"
	"github.com/digitalcircle-com-br/restdb/pkg/util"
	"github.com/gorilla/mux"
)

type Req struct {
	Val        string `json:"val"`
	Expiration int    `json:"expiration"`
}
type Res struct {
	Str  string   `json:"str"`
	Strs []string `json:"strs"`
	Int  int64    `json:"int"`
}

func Setup() error {

	rt := router.Mux()

	rt.Path("/keys/{k}").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k := mux.Vars(r)["k"]
		// req := Req{}
		// err := util.Decode(&req, r)
		// err = util.HttpErr(err, w)
		// if err != nil {
		// 	return
		// }

		cmd := db.Cli().Keys(r.Context(), k)
		if err := util.HttpErr(cmd.Err(), w); err != nil {
			return
		}
		str, err := cmd.Result()
		if err = util.HttpErr(err, w); err != nil {
			return
		}
		res := Res{}
		res.Strs = str
		util.Encode(&res, w)
	})

	rt.Path("/get/{k}").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k := mux.Vars(r)["k"]
		// req := Req{}
		// err := util.Decode(&req, r)
		// err = util.HttpErr(err, w)
		// if err != nil {
		// 	return
		// }

		cmd := db.Cli().Get(r.Context(), k)
		if err := util.HttpErr(cmd.Err(), w); err != nil {
			return
		}
		str, err := cmd.Result()
		if err = util.HttpErr(err, w); err != nil {
			return
		}
		res := Res{}
		res.Str = str
		util.Encode(&res, w)
	})

	rt.Path("/set/{k}").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k := mux.Vars(r)["k"]
		req := Req{}
		err := util.Decode(&req, r)
		err = util.HttpErr(err, w)
		if err != nil {
			return
		}

		cmd := db.Cli().Set(r.Context(), k, req.Val, time.Duration(req.Expiration)*time.Second)
		if err = util.HttpErr(cmd.Err(), w); err != nil {
			return
		}
		str, err := cmd.Result()
		if err = util.HttpErr(err, w); err != nil {
			return
		}
		res := Res{Str: str}
		util.Encode(&res, w)
	})

	rt.Path("/del/{k}").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		k := mux.Vars(r)["k"]
		
		cmd := db.Cli().Del(r.Context(), k)
		if err := util.HttpErr(cmd.Err(), w); err != nil {
			return
		}
		cmdres, err := cmd.Result()
		if err = util.HttpErr(err, w); err != nil {
			return
		}
		res := Res{Int: cmdres}
		util.Encode(&res, w)

	})

	return nil
}
