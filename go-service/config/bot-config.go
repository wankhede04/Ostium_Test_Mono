package config

import "github.com/spf13/viper"

type BotConfig struct {
	RPC        string
	Contract   string
	PrivateKey string
}

func ReadBotConfig() BotConfig {
	return BotConfig{
		RPC:        viper.GetString("watcher.rpc"),
		Contract:   viper.GetString("watcher.contract"),
		PrivateKey: viper.GetString("watcher.private_key"),
	}
}
