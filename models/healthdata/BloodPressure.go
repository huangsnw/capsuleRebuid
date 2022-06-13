package healthdata

import "time"

type BloodPressureDTO struct {
	UserId     string    `bson:"userId"`
	Hret       int       `bson:"hret"`
	Lret       int       `bson:"lret"`
	Heart      int       `bson:"heart"`
	CreateDate time.Time `bson:"createDate"`
}
