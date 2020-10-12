package db

import (
	"context"
	"UnityGame_ABMServer/config"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mgo_client *mongo.Client

func StartInit() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Conf.MongoUrl))
	if err != nil {
		//log
	}
	mgo_client = client

}

func GetClient() *mongo.Client {
	if mgo_client == nil {
		StartInit()
	}
	return mgo_client

}

func Get_Collection(datatable_name string) *mongo.Collection {

	c := GetClient()
	var dt = c.Database(config.Conf.DBName).Collection(datatable_name)
	return dt
}

//
func Get_Collection_RptDB(datatable_name string) *mongo.Collection {

	c := GetClient()
	var dt = c.Database("fbads_android_data").Collection(datatable_name)
	return dt
}