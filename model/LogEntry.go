package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogEntry struct {
	ID        primitive.ObjectID `json:"-" bson:"_id,omitempty" `
	CarrierId int64              `json:"carrierId" bson:"carrier_id,omitempty"`
}

var LogEntryDB = &DBInstance{
	ColName: "logs",
}

func InitLogEntryModel(s *mongo.Database) {
	LogEntryDB.ApplyDatabase(s)
}
