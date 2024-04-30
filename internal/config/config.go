package config

import "github.com/subosito/gotenv"

func Load(path string) error {
	err := gotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}

type BotConfig interface {
	Token() string
	Host() string
}
