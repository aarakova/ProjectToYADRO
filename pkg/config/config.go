package config

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

func NewConfig() *Config {
	var config Config
	var configN NormalizerConfig
	var configC ClientConfig
	flag.StringVar(&configN.InputText,
		"src",
		"",
		"enter your string for analyzer",
	)
	flag.StringVar(&configN.StopWordsFilePath,
		"stop",
		"cmd/task1/stopWords.txt",
		"file containing stop words",
	)
	flag.IntVar(&configC.MaxOffset,
		"n",
		-1,
		"max comics")
	flag.BoolVar(&config.ShowFlag,
		"show",
		true,
		"need show",
	)
	flag.Parse()

	log.SetFormatter(&easy.Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		LogFormat:       "[%lvl%]: %time% - %msg%\n",
	})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)

	log.Info("config was init successful")

	config.NormalizerConfig = &configN
	config.ClientConfig = &configC
	return &config
}

type Config struct {
	*NormalizerConfig
	*ClientConfig

	ShowFlag bool
}

type ClientConfig struct {
	MaxOffset int
}

type NormalizerConfig struct {
	InputText         string
	StopWordsFilePath string
}
