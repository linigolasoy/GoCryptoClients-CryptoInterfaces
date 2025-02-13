package CryptoInterfaces

const ExchangeTypes 
{
	BitUnixFutures = "BitUnixFutures"
}

type IFuturesExchange interface {

	ExchangeType() string;
	Symbols() []IFuturesSymbol;
}
