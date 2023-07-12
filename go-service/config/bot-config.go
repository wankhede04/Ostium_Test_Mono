package config

import "github.com/spf13/viper"

type BotConfig struct {
	RPC                  string
	Contract             string
	PrivateKeyLiquidator string
	PublicKeyLiquidator  string
}

func ReadBotConfig() BotConfig {
	return BotConfig{
		RPC:                  viper.GetString("watcher.rpc"),
		Contract:             viper.GetString("watcher.contract"),
		PrivateKeyLiquidator: viper.GetString("watcher.private_key_liquidator"),
		PublicKeyLiquidator:  viper.GetString("watcher.public_key_liquidator"),
	}
}
