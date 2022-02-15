package spreadcore

import "fmt"

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
	// market.MapPersonNums["user"] = 0
	market.MapPersonNums["lost"] = 0

	return market, nil
}

func (market *Market) OnDay(curday int, nums int) bool {
	if nums > market.MapPersonNums["none"] {
		nums = market.MapPersonNums["none"]
	}

	if nums <= 0 {
		return true
	}

	// newnums := int(float32(market.ConfigSpread.GetRealRetentionRate(1)) / float32(RATE_VALUE) * float32(nums))

	// if newnums == 0 {
	// 	return true
	// }

	market.MapPersonNums["none"] -= nums
	market.MapPersonNums[fmt.Sprintf("user_%v", curday)] = nums

	return false
}

func (market *Market) CountUsers(curday int) (int, int) {
	total := 0
	users := 0
	for cd := 1; cd <= curday; cd++ {
		total += market.MapPersonNums[fmt.Sprintf("user_%v", cd)]

		if cd < curday {
			users += int(float32(market.ConfigSpread.GetRealRetentionRate(curday-cd)) / float32(RATE_VALUE) * float32(market.MapPersonNums[fmt.Sprintf("user_%v", cd)]))
		} else {
			users += market.MapPersonNums[fmt.Sprintf("user_%v", cd)]
		}
	}

	return total, users
}
