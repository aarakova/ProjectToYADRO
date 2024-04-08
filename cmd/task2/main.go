package main

import (
	log "github.com/sirupsen/logrus"

	"project_yadro_2024/pkg/commicsClient"
	"project_yadro_2024/pkg/config"
	"project_yadro_2024/pkg/normalizer"
)

func main() {
	conf := config.NewConfig()
	formatter, err := normalizer.NewFormatter(conf.StopWordsFilePath)
	if err != nil {
		log.Fatal(err)
	}

	client := commicsclient.NewClient("https://xkcd.com", conf.MaxOffset, formatter)

	if conf.ShowFlag {
		log.Println(client.ConvertAll())
	}
}
