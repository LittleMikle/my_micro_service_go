package main

import (
	"github.com/LittleMikle/my_micro_service_go.git/postgres"
	"github.com/LittleMikle/my_micro_service_go.git/rabbit"
	"github.com/rs/zerolog/log"
)

func main() {
	_, err := postgres.ConnectToPostgres()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	err = rabbit.ConnectToRabbit()
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
