package utils

import (
	"capsuleRebuild/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func GetTotalDuraingTime(userid, collectionName string, begin, end time.Time) int64 {
	filter := bson.D{{Key: "userId", Value: userid}, {Key: "createDate", Value: bson.D{{Key: "$gte", Value: begin}, {Key: "$lt", Value: end}}}}
	c, _ := config.GetMongoDatabase().Collection(collectionName).CountDocuments(context.TODO(), filter)
	return c
}
