package main

import (
	"CMS_PUBSUB_INTEGRATION/Services"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

const port = ":7070"

func main() {

	fmt.Println("Application started in :", time.Now())
	server_config := &fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
		Concurrency:   1024 * 1024,
		AppName:       "PREAPID_CARD",
		ReadTimeout:   100 * time.Second,
		WriteTimeout:  100 * time.Second,
		IdleTimeout:   3 * time.Minute,
	}
	wallet_app := fiber.New(*server_config)
	//wallet_app.Get("/hbtChk", Services.HeartBeat)

	wallet_app.Post("/postgre_operation", Services.PostgresOperation)
	wallet_app.Post("/txn_bigquery_insert", Services.BigquerOperation)
	wallet_app.Post("/usercreation_bigquery_insert", Services.BigquerOperationTwo)

	fmt.Println("Server ready to listen and serve on port", port)
	server_err := wallet_app.Listen(port)
	if server_err != nil {
		log.Fatalf("Failed to listen and serve on port %s due to error:%v\n", port, server_err)
	}
}
