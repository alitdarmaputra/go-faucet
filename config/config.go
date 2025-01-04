package config

import "os"

type Config struct {
	BotToken         string
	WalletPrivateKey string
	RpcUrl           string
}

func GetConfig() *Config {
	return &Config{
		BotToken:         os.Getenv("BOT_TOKEN"),
		WalletPrivateKey: os.Getenv("WALLET_PRIVATE_KEY"),
		RpcUrl:           os.Getenv("RPC_URL"),
	}
}
