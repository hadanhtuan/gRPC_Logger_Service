package model

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type LogEntry struct {
	// base info
	VersionNo       string     `json:"versionNo,omitempty" bson:"version_no,omitempty"`
	CreatedBy       string     `json:"createdBy,omitempty" bson:"created_by,omitempty"`
	CreatedTime     *time.Time `json:"createdTime,omitempty" bson:"created_time,omitempty"`
	LastUpdatedBy   string     `json:"lastUpdatedBy,omitempty" bson:"last_updated_by,omitempty"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime,omitempty" bson:"last_updated_time,omitempty"`

	ID primitive.ObjectID `json:"-" bson:"_id,omitempty" `

	Name string `json:"name" bson:"name,omitempty"`
	Data string `json:"data" bson:"data,omitempty"`
}

var LogEntryDB = &DBInstance{
	ColName:        "log_entry",
	TemplateObject: &LogEntry{},
}

func InitLogEntryModel(s *mongo.Database) {
	LogEntryDB.ApplyDatabase(s)
	fmt.Println("Connected to log_entry DB")

}
