package config

import "github.com/spf13/viper"

type WatcherConfig struct {
	RPC         string
	Contract    string
	BlockNumber int64
}

func ReadWatcherConfig() WatcherConfig {
	return WatcherConfig{
		RPC:         viper.GetString("watcher.rpc"),
		Contract:    viper.GetString("watcher.contract"),
		BlockNumber: viper.GetInt64("watcher.block_number"),
	}
}
