package models

import (
	"errors"

	"github.com/astaxie/beego/validation"
)

const (
	EventJoin    = "join"
	EventLeave   = "leave"
	EventMessage = "message"
	EventDelete  = "delete"
	EventEdit    = "edit"
	EventFirst   = "first"
)

type Event struct {
	Type      EventType
	User      int64
	MessageId int
	Timestamp int
	Content   string
}

func IfThenElse(condition bool, a interface{}, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

func IsValid(model interface{}) (err error) {
	valid := validation.Validation{}
	b, err := valid.Valid(model)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
			// return errors.New(fmt.Sprintf("%s: %s", err.Key, err.Message))
		}
	}
	return nil
}
