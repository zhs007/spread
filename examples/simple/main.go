package main

import (
	"fmt"

	goutils "github.com/zhs007/goutils"
	spreadcore "github.com/zhs007/spread/core"
	"go.uber.org/zap"
)

func main() {
	cfgSpread := &spreadcore.SpreadConfig{
		SpreadingRate:    5,
		MapRetentionRate: make(map[int]int),
	}

	cfgSpread.MapRetentionRate[1] = 4000
	cfgSpread.MapRetentionRate[30] = 1000
	cfgSpread.MapRetentionRate[99] = 1000

	cfgMarket := &spreadcore.MarketConfig{
		TotalNums: 100000000,
	}

	market, err := spreadcore.NewMarket(cfgSpread, cfgMarket)
	if err != nil {
		goutils.Error("NewMarket",
			zap.Error(err))

		return
	}

	day := 1
	for !market.OnDay(day, 10000) {
		t, cu := market.CountUsers(day)
		fmt.Printf("%v day: none-%v users-%v lastusers-%v\n", day, market.MapPersonNums["none"], float32(t)/100000000.0, float32(cu)/100000000.0)

		day++
	}
}
