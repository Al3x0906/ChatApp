package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Chat struct {
	Id              int64     `json:"chatid"`
	User1           int64     `json:"user1"`
	User2           int64     `json:"user2"`
	Lastmessage     int64     `json:"lastmessage"`
	Lastmessagetime time.Time `json:"lastmessagetime"`
}

func IsAllowed(chat, user int64) bool {
	var result bson.D
	_ = ChatCollection.FindOne(ctx, bson.D{{"id", chat}}).Decode(&result)
	res := result.Map()

	if res["user1"].(int64) == user || res["user2"].(int64) == user {
		return true
	} else {
		return false
	}
}
