package config

import (
	"os"
	"strconv"
)

type AppConf struct {
	Environment string
	Name        string
}

type HttpConf struct {
	Port       string
	XRequestID string
	Timeout    int
}

type SqlDbConf struct {
	Host                   string
	Username               string
	Password               string
	Name                   string
	Port                   string
	SSLMode                string
	Schema                 string
	MaxOpenConn            int
	MaxIdleConn            int
	MaxIdleTimeConnSeconds int64
	MaxLifeTimeConnSeconds int64
}

type MongoDbConf struct {
	Dsn string
}

type LogConf struct {
	Name string
}

// Config ...
type Config struct {
	App     AppConf
	SqlDb   SqlDbConf
	MongoDb MongoDbConf
	Http    HttpConf
	Log     LogConf
}

// NewConfig ...
func Make() Config {
	app := AppConf{
		Environment: os.Getenv("APP_ENV"),
		Name:        os.Getenv("APP_NAME"),
	}

	sqldb := SqlDbConf{
		Host:     os.Getenv("DB_HOST"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	mongodb := MongoDbConf{
		Dsn: os.Getenv("MONGO_DSN"),
	}

	http := HttpConf{
		Port:       os.Getenv("HTTP_PORT"),
		XRequestID: os.Getenv("HTTP_REQUEST_ID"),
	}

	log := LogConf{
		Name: os.Getenv("LOG_NAME"),
	}

	// set default env to local
	if app.Environment == "" {
		app.Environment = "LOCAL"
	}

	// set default port for HTTP
	if http.Port == "" {
		http.Port = "8080"
	}

	httpTimeout, err := strconv.Atoi(os.Getenv("HTTP_TIMEOUT"))
	if err == nil {
		http.Timeout = httpTimeout
	}

	dBMaxOpenConn, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN_CONN"))
	if err == nil {
		sqldb.MaxOpenConn = dBMaxOpenConn
	}

	dBMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err == nil {
		sqldb.MaxIdleConn = dBMaxIdleConn
	}

	dBMaxIdleTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxIdleTimeConnSeconds = int64(dBMaxIdleTimeConnSeconds)
	}

	dBMaxLifeTimeConnSeconds, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE_TIME_CONN_SECONDS"))
	if err == nil {
		sqldb.MaxLifeTimeConnSeconds = int64(dBMaxLifeTimeConnSeconds)
	}

	config := Config{
		App:     app,
		SqlDb:   sqldb,
		MongoDb: mongodb,
		Http:    http,
		Log:     log,
	}

	return config
}
