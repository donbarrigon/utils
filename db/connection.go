package db

import (
	"context"
	"time"

	"github.com/donbarrigon/utils/config"
	"github.com/donbarrigon/utils/herror"
	"github.com/donbarrigon/utils/logs"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

// var Mongo *mongo.Database

func Col(col string) *mongo.Collection {
	return DB.Collection(col)
}

func InitMongoDB() herror.Error {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(config.DbConnectionString).SetServerAPIOptions(serverAPI)
	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(5)
	clientOptions.SetRetryWrites(true)
	clientOptions.SetTimeout(30 * time.Second)

	var e error
	Client, e = mongo.Connect(clientOptions)
	if e != nil {
		logs.Info("🔴💥 Fail to connect db %s: %s", config.DbName, config.DbConnectionString)
		return herror.InternalServerError(e)
	}
	DB = Client.Database(config.DbName)

	logs.Info("🍃 Successful connection to %s: %s", config.DbName, config.DbConnectionString)
	return nil
}

func CloseMongoDB() herror.Error {
	if Client == nil {
		return nil
	}
	if e := Client.Disconnect(context.TODO()); e != nil {
		return herror.InternalServerError(e)
	}
	return nil
}
