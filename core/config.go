package spreadcore

import (
	"math/rand"

	"github.com/zhs007/goutils"
	"go.uber.org/zap"
)

type SpreadConfig struct {
	SpreadingRate    int         // 传播率，万分数
	MapRetentionRate map[int]int // 留存率，万分数；日留存率，至少有2个节点，中间线性插值；一般来说，1表示次留，99表示无限远
	OffRetentionRate int         // 留存率的误差，万分数，+-，是基于当前留存率的比例，如果是500，表示[95%, 105%]
}

func (cfg *SpreadConfig) Check() error {
	if cfg.SpreadingRate < 0 || cfg.SpreadingRate > RATE_VALUE {
		goutils.Error("SpreadConfig.Check:SpreadingRate",
			zap.Int("val", cfg.SpreadingRate),
			zap.Error(ErrInvalidRate))

		return ErrInvalidRate
	}

	if cfg.OffRetentionRate < 0 || cfg.OffRetentionRate > RATE_VALUE {
		goutils.Error("SpreadConfig.Check:OffRetentionRate",
			zap.Int("val", cfg.OffRetentionRate),
			zap.Error(ErrInvalidRate))

		return ErrInvalidRate
	}

	v1, isok := cfg.MapRetentionRate[1]
	if !isok {
		goutils.Error("SpreadConfig.Check:MapRetentionRate[1]",
			zap.Error(ErrInvalidMapRetentionRate))

		return ErrInvalidMapRetentionRate
	}

	if v1 < 0 || v1 > RATE_VALUE {
		goutils.Error("SpreadConfig.Check:MapRetentionRate[1]",
			zap.Int("val", v1),
			zap.Error(ErrInvalidRate))

		return ErrInvalidRate
	}

	v99, isok := cfg.MapRetentionRate[99]
	if !isok {
		goutils.Error("SpreadConfig.Check:MapRetentionRate[99]",
			zap.Error(ErrInvalidMapRetentionRate))

		return ErrInvalidMapRetentionRate
	}

	if v99 < 0 || v99 > RATE_VALUE {
		goutils.Error("SpreadConfig.Check:MapRetentionRate[1]",
			zap.Int("val", v99),
			zap.Error(ErrInvalidRate))

		return ErrInvalidRate
	}

	return nil
}

func (cfg *SpreadConfig) GetRetentionRate(day int) int {
	if day <= 0 {
		return RATE_VALUE
	}

	if day >= 99 {
		return cfg.MapRetentionRate[99]
	}

	v, isok := cfg.MapRetentionRate[day]
	if isok {
		return v
	}

	preday := 1
	nextday := 99
	for k := range cfg.MapRetentionRate {
		if k < day {
			if preday < k {
				preday = k
			}
		} else {
			if nextday > k {
				nextday = k
			}
		}
	}

	return cfg.MapRetentionRate[preday] +
		int(float32(day-preday)*float32(cfg.MapRetentionRate[nextday]-cfg.MapRetentionRate[preday])/float32(nextday-preday))
}

func (cfg *SpreadConfig) GetRealRetentionRate(day int) int {
	if cfg.OffRetentionRate == 0 {
		return cfg.GetRetentionRate(day)
	}

	cr := (rand.Int() % (cfg.OffRetentionRate * 2)) - cfg.OffRetentionRate
	co := 1.0 + float32(cr)/float32(RATE_VALUE)

	return int(float32(cfg.GetRetentionRate(day)) * co)
}

type MarketConfig struct {
	TotalNums int
}
