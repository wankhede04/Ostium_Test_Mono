package config

import "github.com/spf13/viper"

type WatcherConfig struct {
	RPC      string
	Contract string
}

func ReadWatcherConfig() WatcherConfig {
	return WatcherConfig{
		RPC:      viper.GetString("watcher.rpc"),
		Contract: viper.GetString("watcher.contract"),
	}
}
