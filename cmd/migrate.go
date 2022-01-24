package main

import (
	mslog "github.com/Sinbad-B2B-Platform/ms-go-utils/instance/log"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/config"
	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/persistence/postgres"
)

func main() {
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

	logger := mslog.NewLogInstance(
		mslog.LogName(conf.Log.Name),
		mslog.IsProduction(isProd),
		mslog.LogAdditionalFields(m))

	db := postgres.New(conf.SqlDb, logger)

	logger.Printf("DB connection %s SUCCESS", db.DB.Name())

	err := postgres.Migrate(db.DB)
	if err != nil {
		logger.Fatalf("migration err: %s", err)
	}

	logger.Println("Success migration")
}
