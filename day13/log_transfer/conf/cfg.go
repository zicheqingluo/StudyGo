package conf

// LogTransfer 全局配置
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg	`ini:"es"`
}

//kafkacfg ...
type KafkaCfg struct{
	Address	string 	`ini:"address"`
	Topic	string 	`ini:"topic"`
}

//ESCfg ...
type ESCfg	struct{
	Address	string 	`ini:"address"`
	ChanSize	int	`ini::"chansize"`
}