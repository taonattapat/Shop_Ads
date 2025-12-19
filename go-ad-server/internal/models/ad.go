package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Ad struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title     string             `bson:"title" json:"title"`
	ImageURL  string             `bson:"image_url" json:"image_url"`
	TargetURL string             `bson:"target_url" json:"target_url"`
	Budget    float64            `bson:"budget" json:"budget"`
	Spent     float64            `bson:"spent" json:"spent"`
	Status    string             `bson:"status" json:"status"` // "active", "paused", "finished"
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}
