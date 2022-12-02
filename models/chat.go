package models

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func IsAllowed(id, user int64) (*Chat, error) {
	chat := Chat{}
	err := chat.Get(id)
	if err != nil {
		return nil, err
	}

	if chat.User1 == user || chat.User2 == user {
		return &chat, nil
	} else {
		return nil, errors.New("user not allowed")
	}
}

type Chat struct {
	Id              int64     `json:"chatid"`
	User1           int64     `json:"user1"`
	User2           int64     `json:"user2"`
	Lastmessage     int64     `json:"lastmessage"`
	Lastmessagetime time.Time `json:"lastmessagetime"`
}

func (chat Chat) Get(id int64) error {
	var result bson.D
	q := ChatCollection.FindOne(ctx, bson.D{{"id", id}})
	err := q.Decode(&result)

	if err != nil {
		fmt.Println(result)
		fmt.Println("goa", err.Error())
		return err
	}
	res := result.Map()

	chat.Id = res["id"].(int64)
	chat.User1 = res["user1"].(int64)
	chat.User2 = res["user2"].(int64)
	chat.Lastmessage = res["lastmessage"].(int64)
	chat.Lastmessagetime = res["lastmessagetime"].(time.Time)

	return nil
}

func (chat Chat) objToBson(id int64) bson.D {
	var result bson.D
	_ = counter.FindOne(ctx, bson.D{{"collection", "Chat"}}).Decode(&result)
	ChatCounter = result.Map()["count"].(int64)

	id = IfThenElse(id == -1, ChatCounter, id).(int64)
	return bson.D{
		{"id", id},
		{"user1", chat.User1},
		{"user2", chat.User2},
		{"lastmessage", chat.Lastmessage},
		{"lastmessagetime", chat.Lastmessagetime},
	}
}
