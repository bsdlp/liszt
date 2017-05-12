package main

import (
	"log"
	"net/http"

	"github.com/kelseyhightower/envconfig"
	"github.com/liszt-code/liszt/pkg/registry"
	"github.com/liszt-code/liszt/pkg/registry/service"
)

func main() {
	cfg := new(registry.Config)
	err := envconfig.Process("liszt", cfg)
	if err != nil {
		log.Fatalln(err)
	}
	registrar, err := registry.New(cfg)
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
