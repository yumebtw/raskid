package config

import (
	"errors"
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Bot      Bot      `yaml:"bot"`
	Database Database `yaml:"database"`
}

type Bot struct {
	Token string `yaml:"token"`
}

type Database struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

var Cfg Config

func MustLoad() *Config {
	var filepath string

	flag.StringVar(&filepath, "config", "", "path to config file")
	flag.Parse()

	if filepath == "" {
		panic(errors.New("config file path is not provided"))
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		panic(errors.New("config file does not exist"))
	}

	var cfg Config
	if err := cleanenv.ReadConfig(filepath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}

	return &cfg
}
