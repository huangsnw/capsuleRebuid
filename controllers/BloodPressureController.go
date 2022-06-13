package controllers

import (
	"capsuleRebuild/config"
	"capsuleRebuild/models/healthdata"
	"capsuleRebuild/models/res"
	"capsuleRebuild/utils"
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DetailBloodPressure(ctx *gin.Context) {
	userid := ctx.Query("userId")
	begin := utils.FormatString2Time(ctx.Query("begin"))
	end := utils.FormatString2Time(ctx.Query("end"))
	end = utils.NextDay(end)

	collection := config.GetMongoDatabase().Collection("bloodPressureDTO")
	filter := bson.D{{Key: "userId", Value: userid}, {Key: "createDate", Value: bson.D{{Key: "$gte", Value: begin}, {Key: "$lt", Value: end}}}}

	findOptions := options.Find()
	findOptions.SetLimit(10)
	findOptions.SetSort(bson.M{"createDate": -1})

	result := []healthdata.BloodPressureDTO{}

	c, _ := collection.Find(context.TODO(), filter, findOptions)
	for c.Next(context.TODO()) {
		var bp healthdata.BloodPressureDTO
		c.Decode(&bp)
		result = append(result, bp)
	}
	defer c.Close(context.TODO())
	r := res.Result{}
	ctx.JSON(200, r.SUCESS(result))
}
