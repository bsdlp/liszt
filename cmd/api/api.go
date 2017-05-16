package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/liszt-code/liszt/pkg/registry/rdbms"
	"github.com/liszt-code/liszt/pkg/registry/service"
)

func main() {
	cfg := new(rdbms.Config)
	err := envconfig.Process("liszt", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	registrar, err := rdbms.New(cfg)
	if err != nil {
		log.Fatalln(err)
	}
	server := &http.Server{
		Addr: ":8080",
		Handler: &service.Service{
			Registrar: registrar,
		},
	}
	log.Fatalln(server.ListenAndServe())
}
