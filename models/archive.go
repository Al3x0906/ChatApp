package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
	"time"
)

func connect() (*mongo.Client, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 999999*time.Hour)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
	if err == nil {
		fmt.Println("Mongo Connected")
	} else {
		fmt.Println(err.Error())
	}
	return client, ctx, cancel
}

var (
	Client, ctx, _     = connect()
	DataBase           = Client.Database("chatapp")
	MessagesCollection = DataBase.Collection("messages")
	ChatCollection     = DataBase.Collection("chat")
	counter            = DataBase.Collection("counter")
	MessageCounter     int64
	ChatCounter        int64
	MessageMutex       sync.Mutex
	ZeroTime           = time.Time{}
)

type EventType string

func init() {
	var result bson.D
	_ = counter.FindOne(ctx, bson.D{{"collection", "Messages"}}).Decode(&result)
	MessageCounter = result.Map()["count"].(int64)

	result = bson.D{}
	_ = counter.FindOne(ctx, bson.D{{"collection", "Chat"}}).Decode(&result)
	ChatCounter = result.Map()["count"].(int64)
}
