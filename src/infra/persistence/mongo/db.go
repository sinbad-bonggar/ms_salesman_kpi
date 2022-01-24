package mongo

import (
	"context"

	"github.com/sinbad-bonggar/ms_salesman_kpi/src/infra/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoDb struct {
	DB *mongo.Client
}

func New(conf config.MongoDbConf, logger *logrus.Logger) *mongoDb {
	dsn := conf.Dsn

	clientOptions := options.Client()
	clientOptions.ApplyURI(dsn)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic("Failed to connect to database!")
	}

	logger.Println("mongo database connection success")

	return &mongoDb{DB: client}
}
