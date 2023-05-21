package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var UsersCollection *mongo.Collection
var OrdersCollection *mongo.Collection

func InitMongoDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	MongoClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		println("mongodb 连接失败")
		panic(err)
	}

	err = MongoClient.Ping(ctx, nil)
	if err == nil {
		println("mongodb 连接成功")
	}

	// 选择数据库home
	database := MongoClient.Database("home")
	// users表
	UsersCollection = database.Collection("users")
	// orders表
	OrdersCollection = database.Collection("orders")
	return MongoClient
}
