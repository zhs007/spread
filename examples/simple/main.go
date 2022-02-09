package main

import (
	"fmt"

	goutils "github.com/zhs007/goutils"
	spreadcore "github.com/zhs007/spread/core"
	"go.uber.org/zap"
)

func main() {
	cfgSpread := &spreadcore.SpreadConfig{
		SpreadingRate: 5,
		RetentionRate: 600,
	}

	cfgMarket := &spreadcore.MarketConfig{
		TotalNums: 100000000,
	}

	market, err := spreadcore.NewMarket(cfgSpread, cfgMarket)
	if err != nil {
		goutils.Error("NewMarket",
			zap.Error(err))

		return
	}

	day := 0
	for !market.OnDay(10000) {
		fmt.Printf("%v day: none-%v users-%v\n", day, market.MapPersonNums["none"], market.MapPersonNums["user"])

		day++
	}
}
