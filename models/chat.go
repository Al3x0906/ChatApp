package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func IsAllowed(id, user int64) (*Chat, error) {
	chat := &Chat{}
	err := chat.Get(id)
	if err != nil {
		return nil, err
	}

	if chat.User1 == user || chat.User2 == user {
		return chat, nil
	} else {
		log.Println("fuck", chat, user)
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

func (chat *Chat) Get(id int64) error {
	var res bson.M
	q := ChatCollection.FindOne(ctx, bson.D{{"id", id}})
	err := q.Decode(&res)

	if err != nil {
		return err
	}

	chat.Id = res["id"].(int64)
	chat.User1 = res["user1"].(int64)
	chat.User2 = res["user2"].(int64)
	chat.Lastmessage = res["lastmessage"].(int64)
	if chat.Lastmessage == -1 {
		chat.Lastmessagetime = time.Time{}
	} else {
		chat.Lastmessagetime = res["lastmessagetime"].(time.Time)
	}

	return nil
}

func (chat *Chat) objToBson() bson.D {
	return bson.D{
		{"id", chat.Id},
		{"user1", chat.User1},
		{"user2", chat.User2},
		{"lastmessage", chat.Lastmessage},
		{"lastmessagetime", chat.Lastmessagetime},
	}
}

// TODO: Add new chat method and use mutex
//	id = IfThenElse(id == -1, ChatCounter, id).(int64)
