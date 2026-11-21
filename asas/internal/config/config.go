package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	General General `toml:"general"`
	Bar     Bar     `toml:"bar"`
	Clock   Clock   `toml:"clock"`
}

type General struct {
	Theme    string `toml:"theme"`
	LogLevel string `toml:"log_level"`
}
type Bar struct {
	Position string `toml:"position"`
	Shape    string `toml:"shape"`
	Height   int    `toml:"height"`
}
type Clock struct {
	Format     string `toml:"format"`
	DateFormat string `toml:"date_format"`
}

func Load(path string) (Config, error) {
	var c Config
	_, err := toml.DecodeFile(path, &c)
	return c, err
}

func Defaults() Config {
	return Config{
		General: General{Theme: "default", LogLevel: "info"},
		Bar:     Bar{Position: "top", Shape: "islands", Height: 40},
		Clock:   Clock{Format: "%H:%M", DateFormat: "%a %d %b"},
	}
}