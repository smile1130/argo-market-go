package main

import (
	"github.com/gocraft/web"

	"argomarket/payment-gateway/cryptocurrencies"
	"argomarket/payment-gateway/cryptocurrencies/bitcoin"
	"argomarket/payment-gateway/cryptocurrencies/bitcoincash"
	"argomarket/payment-gateway/cryptocurrencies/ethereum"

	"argomarket/payment-gateway/exchange"
)

func ConfigureRouter(router *web.Router) *web.Router {

	bitcoinRouter := router.Subrouter(Context{}, "/bitcoin")
	bitcoin.ConfigureRouter(bitcoinRouter)

	bitcoinCashRouter := router.Subrouter(Context{}, "/bitcoin_cash")
	bitcoincash.ConfigureRouter(bitcoinCashRouter)

	ethereumRouter := router.Subrouter(Context{}, "/ethereum")
	ethereum.ConfigureRouter(ethereumRouter)

	// Exchange
	exchangeRouter := router.Subrouter(Context{}, "/exchange")
	exchange.ConfigureRouter(exchangeRouter)

	// Currency
	router.Get("/currency/:base_currency", cryptocurrencies.ViewShowCurrency)
	return router
}
