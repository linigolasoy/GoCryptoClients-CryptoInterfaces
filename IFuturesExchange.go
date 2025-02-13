package CryptoInterfaces

const (
	BitUnixFutures string = "BitUnixFutures"
)

type IFuturesExchange interface {
	ExchangeType() string
	Symbols() []IFuturesSymbol
}
