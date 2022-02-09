package spreadcore

type SpreadConfig struct {
	SpreadingRate int // 传播率，万分数
	RetentionRate int // 留存率，万分数
}

type MarketConfig struct {
	TotalNums int
}
