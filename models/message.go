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
	Content  []byte    `json:"message"`
	Sent     time.Time `json:"sent"`
	Seen     time.Time `json:"seen"`
	Edited   time.Time `json:"edited"`
}

func (message *Message) objToBson(id int64) bson.D {
	MessageMutex.Lock()
	message.Id = IfThenElse(id == -1, MessageCounter, id).(int64)
	MessageCounter += 1
	MessageMutex.Unlock()

	return bson.D{
		{"id", message.Id},
		{"chat", message.Chat},
		{"sender", message.Sender},
		{"Receiver", message.Receiver},
		{"content", message.Content},
		{"sent", message.Sent},
		{"seen", message.Seen},
		{"edited", message.Edited},
	}
}

func (message *Message) Insert() error {
	doc := message.objToBson(-1)

	_, _ = counter.UpdateOne(ctx,
		bson.D{{"collection", "Messages"}},
		bson.D{{"$set", bson.D{{"count", MessageCounter}}}})

	_, err := MessagesCollection.InsertOne(ctx, doc)
	return err
}

func (message *Message) MapMessage(d bson.M) {
	message.Id = d["id"].(int64)
	message.Chat = d["chat"].(int64)
	message.Sender = d["sender"].(int64)
	message.Receiver = d["receiver"].(int64)
	message.Content = d["content"].([]byte)
	message.Sent = d["sent"].(time.Time)
	message.Seen = d["seen"].(time.Time)
	message.Edited = d["edited"].(time.Time)
}

func (message *Message) Get(id, chat, user int64) error {
	filter := bson.D{{"id", id}, {"chat", chat}}

	var d bson.M
	err := MessagesCollection.FindOne(ctx, filter).Decode(&d)

	if err == nil {
		fmt.Println(d)
		if d["sender"].(int64) == user || d["receiver"].(int64) == user {
			message.MapMessage(d)
		} else {
			err = errors.New("user not allowed to see message")
		}
	}

	return err
}

func (message *Message) Update() (int64, error) {
	filter := bson.D{{"id", message.Id}, {"chat", message.Chat}, {"sender", message.Sender}}
	doc := bson.D{{"$set", bson.D{{"content", message.Content}, {"edited", time.Now()}}}}

	result, err := MessagesCollection.UpdateOne(ctx, filter, doc)
	return result.ModifiedCount, err
}

func (message *Message) SetSeen() error {
	filter := bson.D{{"id", message.Id}}
	doc := bson.D{{"$set", bson.D{{"seen", message.Seen}}}}

	_, err := MessagesCollection.UpdateOne(ctx, filter, doc)
	return err
}

func (message *Message) Delete() (int64, error) {
	filter := bson.D{{"id", message.Id}, {"chat", message.Chat}, {"sender", message.Sender}}
	result, err := MessagesCollection.DeleteOne(ctx, filter)

	return result.DeletedCount, err
}

// TODO: Load old or messages on join
