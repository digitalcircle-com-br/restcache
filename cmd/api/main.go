package main

import (
	"log"
	"net/http"

	"github.com/digitalcircle-com-br/buildinfo"
	"github.com/digitalcircle-com-br/restdb/pkg"
	"github.com/digitalcircle-com-br/restdb/pkg/infra/router"
)

func main() {
	log.Println(buildinfo.String())
	err := pkg.Setup()
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Fatal(http.ListenAndServe(":8080", router.Mux()))
}
