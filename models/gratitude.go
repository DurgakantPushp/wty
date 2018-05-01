package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Gratitude represents a gratitude sent by a user
type Gratitude struct {
	ID        bson.ObjectId `bson:"id" json:"id,omitempty"`
	Sender    string        `bson:"sender" json:"sender"`
	Recipient string        `bson:"recipient" json:"recipient"`
	Content   string        `bson:"cnt" json:"cnt"`
}

// NewGratitude returns *Gratitude
func NewGratitude(sender string) *Gratitude {
	return &Gratitude{
		ID:     bson.NewObjectId(),
		Sender: sender,
		// Recipient: recipient,
		// Content:   cnt,
	}
}
