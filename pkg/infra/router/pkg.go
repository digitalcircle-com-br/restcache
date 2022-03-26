package router

import "github.com/gorilla/mux"

var mainMux = mux.NewRouter()

func Setup() error {

	return nil
}

func Mux() *mux.Router {
	return mainMux
}
