package main

import (
	"log"
	"net/http"
	"swift/api/handlers"
	"swift/internal/configs"
	"swift/internal/repositories"
	"swift/internal/services"
	"swift/loggers"
	"swift/pkg/db"

	"go.uber.org/zap"
)

func main() {
	log.Println("Starting logging...")
	logger, err := loggers.InitLogger()

	if err != nil {
		log.Fatal(err)
		return
	}
	// log.Printf("%+v\n",logger)
	// log.Println("Successfully! ))")
	defer func(logg *zap.Logger) {
		//записывает все события в logger
		err := logg.Sync()
		if err != nil {
			log.Println(err)
		}
	}(logger)

	//get configs
	// log.Println("Start reading from file config...")
	err = configs.InitConfigs("./internal/configs/config.json")
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println("Success! ))")

	//get db
	// log.Println("Open connection with Db")
	database, err := db.OpenConnection()
	if err != nil {
		log.Fatalln("ERROR while connection with database")
		return
	}
	// log.Printf("%+v\n", database)
	// log.Println("Success! ))")
	defer db.CloseConnection()

	repository := repositories.NewRepository(database)
	// log.Printf("%+v\n",repository)
	service := services.NewService(repository)
	// log.Printf("%+v\n",service)
	handler := handlers.NewHandler(service, logger)
	// log.Printf("%+v\n",handler)

	router := handlers.InitRouter(handler)
	// log.Printf("%+v\n",router)
	// log.Println("Connection with routers was successfully")
	srv := http.Server{
		Addr:    configs.Settings.Server.Host + configs.Settings.Server.Port,
		Handler: router,
	}

	// log.Printf("%+v\n",srv)
	// log.Printf("%+v\n",srv.Addr)
	// log.Printf("%+v\n",srv.Handler)
	log.Println("Start!!")
	log.Println(srv.Addr)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
