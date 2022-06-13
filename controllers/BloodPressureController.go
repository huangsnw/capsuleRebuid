package controllers

import (
	"capsuleRebuild/config"
	"capsuleRebuild/models/healthdata"
	"capsuleRebuild/models/pages"
	"capsuleRebuild/models/res"
	"capsuleRebuild/utils"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DetailBloodPressure(ctx *gin.Context) {
	userid := ctx.Query("userId")
	pageIndex, _ := strconv.Atoi(ctx.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	begin := utils.FormatString2Time(ctx.Query("begin"))
	end := utils.FormatString2Time(ctx.Query("end"))
	end = utils.NextDay(end)
	// query data
	filter := bson.D{{Key: "userId", Value: userid}, {Key: "createDate", Value: bson.D{{Key: "$gte", Value: begin}, {Key: "$lt", Value: end}}}}
	findOptions := options.Find().SetSort(bson.M{"createDate": -1}).SetSkip((int64(pageIndex) - 1) * int64(pageSize)).SetLimit(int64(pageSize))
	result := []healthdata.BloodPressureDTO{}
	c, _ := config.GetMongoDatabase().Collection("bloodPressureDTO").Find(context.TODO(), filter, findOptions)
	// query total
	total := utils.GetTotalDuraingTime(userid, "bloodPressureDTO", begin, end)
	for c.Next(context.TODO()) {
		var bp healthdata.BloodPressureDTO
		c.Decode(&bp)
		result = append(result, bp)
	}
	defer c.Close(context.TODO())
	page := pages.Page{PageIndex: pageIndex, PageSize: pageSize, Total: total, List: result}
	res := res.Result{}
	ctx.JSON(200, res.SUCESS(page))
}

func LatestBloodPressure(ctx *gin.Context) {
	userid := ctx.Query("userId")
	// query data
	var bloodPressureDTO healthdata.BloodPressureDTO
	filter := bson.D{{Key: "userId", Value: userid}}
	//filter := bson.M{"userId": userid}
	findOptions := options.FindOne().SetSort(bson.M{"createDate": -1})
	config.GetMongoDatabase().Collection("bloodPressureDTO").FindOne(context.TODO(), filter, findOptions).Decode(&bloodPressureDTO)
	r := res.Result{}
	ctx.JSON(200, r.SUCESS(bloodPressureDTO))
}
