package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"

	"argomarket/payment-gateway/settings"
)

type Context struct {
}

func RunServer() {
	router := web.New(Context{})
	ConfigureRouter(router)

	address := fmt.Sprintf("%s:%d", settings.APPLICATION_SETTINGS.Host, settings.APPLICATION_SETTINGS.Port)
	println("Running server on " + address)
	http.ListenAndServe(address, router)
}
