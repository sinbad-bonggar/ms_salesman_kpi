package postgres

import (
	"database/sql"
	"fmt"

	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
)

type postgresDb struct {
	DB    *gorm.DB
	SqlDB *sql.DB
}

func New(conf config.SqlDbConf, logger *logrus.Logger) *postgresDb {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		conf.Host,
		conf.Username,
		conf.Password,
		conf.Name,
		conf.Port,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: log.Default.LogMode(log.Warn),
	})
	if err != nil {
		panic("Failed to connect to database!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.Fatalf("database err: %s", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}

	logger.Printf("sql database connection %s success", db.Name())

	return &postgresDb{
		DB:    db,
		SqlDB: sqlDB,
	}
}
