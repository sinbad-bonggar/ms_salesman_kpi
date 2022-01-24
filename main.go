package main

import (
	"context"
	"database/sql"

	ms_log "github.com/Sinbad-B2B-Platform/ms-go-utils/instance/log"
	_ "github.com/joho/godotenv/autoload"
	usecases "github.com/sinbad-bonggar/ms_salesman_kpi/src/app/use_cases"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/config"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/persistence/postgres"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/interface/rest"
	"github.com/sirupsen/logrus"
)

func main() {
	// init context
	ctx := context.Background()

	// read the server environment variables
	conf := config.Make()

	// check is in production mode
	isProd := false
	if conf.App.Environment == "PRODUCTION" {
		isProd = true
	}

	// logger setup
	m := make(map[string]interface{})
	m["env"] = conf.App.Environment
	m["service"] = conf.App.Name
	logger := ms_log.NewLogInstance(
		ms_log.LogName(conf.Log.Name),
		ms_log.IsProduction(isProd),
		ms_log.LogAdditionalFields(m))

	// open connection to persistence storage
	postgresdb := postgres.New(conf.SqlDb, logger)

	// gracefully close connection to persistence storage
	defer func(l *logrus.Logger, sqlDB *sql.DB, dbName string) {
		err := sqlDB.Close()
		if err != nil {
			l.Errorf("error closing sql database %s: %s", dbName, err)
		} else {
			l.Printf("sql database %s successfuly closed.", dbName)
		}
	}(logger, postgresdb.SqlDB, postgresdb.DB.Name())

	// for using mongo
	// mongodb := mongo.New(conf.MongoDb)

	// Bussiness Domain
	// We're using Domain Driven Design (DDD) Paradigm and Clean Code Architecture

	// Order
	orderRepository := postgres.NewOrderRepository(postgresdb.DB)

	// HTTP Handler
	// the server already implements a graceful shutdown.
	httpServer, err := rest.New(
		conf.Http,
		isProd,
		logger,
		usecases.AllUseCases{
			OrderUseCase: usecases.NewOrderUseCase(orderRepository),
		},
	)
	if err != nil {
		panic(err)
	}
	httpServer.Start(ctx)

	// Kafka Interface
	// _, kProducer, err := infra.RegisterKafka()
	// if err != nil {
	// 	log.Fatal("Unable to create kafka")
	// }
	// defer kProducer.Producer.Close()
}
