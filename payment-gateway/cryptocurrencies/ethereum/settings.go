package ethereum

import (
	"argomarket/payment-gateway/settings"
)

var (
	PAYMENT_GATE_SETTINGS = settings.LoadPaymentGateSettings()
	WEI_IN_ETH            = float64(1e18)
)
