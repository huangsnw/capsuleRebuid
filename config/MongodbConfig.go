package config

import (
	"context"
	"log"
	"strings"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgoDB *mongo.Database

func InitMongoDB() {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("mongodb://")
	strBuilder.WriteString(viper.GetString("mongourls"))
	strBuilder.WriteString(":")
	strBuilder.WriteString(viper.GetString("mongoport"))

	log.Println("connect mongodb on ", strBuilder.String())

	clientOptions := options.Client().ApplyURI(strBuilder.String())
	client, e := mongo.Connect(context.TODO(), clientOptions)
	PrintError(e, "mongodb client error")
	// 自测
	e = client.Ping(context.TODO(), nil)
	PrintError(e, "mongodb client Ping faild")
	mgoDB = client.Database("healthcare")

}

func GetMongoDatabase() *mongo.Database {
	return mgoDB
}
