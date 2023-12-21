package exchange

import (
	btc "argomarket/payment-gateway/cryptocurrencies/bitcoin"
	bch "argomarket/payment-gateway/cryptocurrencies/bitcoincash"
)

// TransferBitcoinCashToBitcoin transfers coins from Bitcoin Cash Wallet to a Bitcoin Wallet
// This is needed because a lot of users mistakenly send BTC to a BCH wallets and then are unable
// to use their funds
func TransferBitcoinCashToBitcoin(bchWalletPK, btcWalletPK string) (btc.BTCPaymentResult, error) {

	bchWallet, err := bch.FindBitcoinCashWalletByPublicKey(bchWalletPK)
	if err != nil {
		return btc.BTCPaymentResult{}, err
	}

	importedBtcWallet, err := btc.ImportBitcoinWallet(bchWallet.Type, bchWallet.PrivateKey)
	if err != nil {
		return btc.BTCPaymentResult{}, err
	}

	payments := []btc.BTCPayment{
		{Address: btcWalletPK, Percent: 1.0},
	}

	return importedBtcWallet.SendWithPercentSplit(payments)
}
