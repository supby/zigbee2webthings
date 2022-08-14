package main

import (
	"log"

	"github.com/supby/zigbee2webthings/api/restapi"
	apiservice "github.com/supby/zigbee2webthings/internal/api-service"
)

func main() {
	var apiConfig restapi.Config

	err := apiConfig.WithDefaultFlags().Parse()
	if err != nil {
		log.Fatal(err)
	}

	svc := &apiservice.WebThingsApiService{}

	api := restapi.NewServer(svc, &apiConfig)

	err = api.RunWithSigHandler()
	if err != nil {
		log.Fatal(err)
	}
}
