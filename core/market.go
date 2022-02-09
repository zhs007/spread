package spreadcore

type Market struct {
	ConfigSpread  *SpreadConfig
	ConfigMarket  *MarketConfig
	MapPersonNums map[string]int
}

func NewMarket(cfgSpread *SpreadConfig, cfgMarket *MarketConfig) (*Market, error) {
	market := &Market{
		ConfigSpread:  cfgSpread,
		ConfigMarket:  cfgMarket,
		MapPersonNums: make(map[string]int),
	}

	market.MapPersonNums["none"] = cfgMarket.TotalNums
	market.MapPersonNums["user"] = 0
	market.MapPersonNums["lost"] = 0

	return market, nil
}

func (market *Market) OnDay(nums int) bool {
	if nums > market.MapPersonNums["none"] {
		nums = market.MapPersonNums["none"]
	}

	if nums <= 0 {
		return true
	}

	newnums := int(float32(market.ConfigSpread.RetentionRate) / float32(RATE_VALUE) * float32(nums))

	if newnums == 0 {
		return true
	}

	market.MapPersonNums["none"] -= newnums
	market.MapPersonNums["user"] += newnums

	return false
}
