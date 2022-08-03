package util

const (
	USD = "USD"
	CAD = "CAD"
	EUR = "EUR"
	CNY = "CNY"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, CAD, EUR, CNY:
		return true
	}
	return false
}
