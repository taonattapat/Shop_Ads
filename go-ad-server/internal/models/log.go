package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TrackingLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AdID      primitive.ObjectID `bson:"ad_id" json:"ad_id"`
	Type      string             `bson:"type" json:"type"` // "view", "click"
	IP        string             `bson:"ip" json:"ip"`
	UserAgent string             `bson:"user_agent" json:"user_agent"`
	Timestamp time.Time          `bson:"timestamp" json:"timestamp"`
}
