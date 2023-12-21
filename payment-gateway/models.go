package main

import (
	"argomarket/payment-gateway/cryptocurrencies/bitcoin"
	"argomarket/payment-gateway/cryptocurrencies/bitcoincash"
	"argomarket/payment-gateway/cryptocurrencies/ethereum"
	"argomarket/payment-gateway/db"
)

func SyncDatabase() {
	db.Database.AutoMigrate(
		&bitcoin.BitcoinWallet{},
		&bitcoin.BitcoinWalletBalance{},

		&bitcoincash.BitcoinCashWallet{},
		&bitcoincash.BitcoinCashWalletBalance{},

		&ethereum.EthereumWallet{},
		&ethereum.EthereumWalletBalance{},
	)

	ethereum.SetupViews()
	bitcoin.SetupViews()
	bitcoincash.SetupViews()
}
