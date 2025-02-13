package CryptoInterfaces

// Symbol interface

type IFuturesSymbol interface {
	Symbol() string
	BaseAsset() string
	QuoteAsset() string
}
