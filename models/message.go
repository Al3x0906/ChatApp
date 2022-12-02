package models

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Message struct {
	Id       int64     `json:"messageid"`
	Chat     int64     `json:"chat"`
	Sender   int64     `json:"sender"`
	Receiver int64     `json:"receiver"`
	Content  string    `json:"message"`
	Sent     time.Time `json:"sent"`
	Seen     time.Time `json:"seen"`
	Edited   time.Time `json:"edited"`
}

func (message Message) objToBson(id int64) bson.D {
	var result bson.D
	_ = counter.FindOne(ctx, bson.D{{"collection", "Messages"}}).Decode(&result)
	MessageCounter = result.Map()["count"].(int64)

	id = IfThenElse(id == -1, MessageCounter, id).(int64)
	return bson.D{
		{"id", id},
		{"chat", message.Chat},
		{"sender", message.Sender},
		{"Receiver", message.Receiver},
		{"content", message.Content},
		{"sent", message.Sent},
		{"seen", message.Seen},
		{"edited", message.Edited},
	}
}

func (message Message) Insert() error {
	message.Sent = time.Now()

	doc := message.objToBson(-1)
	_, err := MessagesCollection.InsertOne(ctx, doc)

	if err == nil {
		_, _ = counter.UpdateOne(ctx,
			bson.D.Map(bson.D{{"MessagesCollection", "Messages"}}),
			bson.D{{"count", MessageCounter + 1}})
	}

	return err
}

func MapMessage(message Message, d bson.M) {
	message.Id = d["id"].(int64)
	message.Chat = d["chat"].(int64)
	message.Sender = d["sender"].(int64)
	message.Receiver = d["receiver"].(int64)
	message.Content = d["content"].(string)
	message.Sent = d["sent"].(time.Time)
	message.Seen = d["seen"].(time.Time)
	message.Edited = d["edited"].(time.Time)
}

func (message Message) Get(id, chat, user int64) error {
	filter := bson.D{{"id", id}, {"chat", chat}}

	var result bson.D
	err := MessagesCollection.FindOne(ctx, filter).Decode(&result)

	if err == nil {
		d := result.Map()
		fmt.Println(d)
		if d["sender"].(int64) == user || d["receiver"].(int64) == user {
			MapMessage(message, d)
		} else {
			err = errors.New("user not allowed to see message")
		}
	}

	return err
}

func (message Message) Update(id, chat int64) error {
	filter := bson.D{{"id", id}, {"chat", chat}}
	doc := message.objToBson(id)

	_, err := MessagesCollection.UpdateOne(ctx, filter, doc)
	return err
}

func (message Message) Delete(id, chat int64) error {
	filter := bson.D{{"id", id}, {"chat", chat}}
	_, err := MessagesCollection.DeleteOne(ctx, filter)

	return err
}

// TODO: Load old or messages on join
