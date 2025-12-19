package repository

import (
	"context"
	"time"

	"go-ad-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LogRepository struct {
	Collection *mongo.Collection
}

func NewLogRepository(db *mongo.Database) *LogRepository {
	return &LogRepository{
		Collection: db.Collection("tracking_logs"),
	}
}

func (r *LogRepository) CreateLog(log *models.TrackingLog) error {
	log.Timestamp = time.Now()
	_, err := r.Collection.InsertOne(context.Background(), log)
	return err
}

func (r *LogRepository) GetLogs(limit int64) ([]models.TrackingLog, error) {
	opts := options.Find().SetSort(bson.M{"timestamp": -1}).SetLimit(limit)
	cursor, err := r.Collection.Find(context.Background(), bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	var logs []models.TrackingLog
	if err = cursor.All(context.Background(), &logs); err != nil {
		return nil, err
	}
	return logs, nil
}

func (r *LogRepository) GetStats(adID string) (map[string]interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(adID)
	if err != nil {
		return nil, err
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"ad_id": objID}}},
		{{Key: "$group", Value: bson.M{
			"_id":   "$type",
			"count": bson.M{"$sum": 1},
		}}},
	}

	cursor, err := r.Collection.Aggregate(context.Background(), pipeline)
	if err != nil {
		return nil, err
	}

	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		return nil, err
	}

	stats := map[string]interface{}{
		"views":  0,
		"clicks": 0,
	}

	for _, res := range results {
		if t, ok := res["_id"].(string); ok {
			if t == "view" {
				stats["views"] = res["count"]
			} else if t == "click" {
				stats["clicks"] = res["count"]
			}
		}
	}

	// Calculate CTR
	var views, clicks float64

	switch v := stats["views"].(type) {
	case int:
		views = float64(v)
	case int32:
		views = float64(v)
	case int64:
		views = float64(v)
	case float64:
		views = v
	}

	switch v := stats["clicks"].(type) {
	case int:
		clicks = float64(v)
	case int32:
		clicks = float64(v)
	case int64:
		clicks = float64(v)
	case float64:
		clicks = v
	}

	ctr := 0.0
	if views > 0 {
		ctr = clicks / views * 100
	}
	stats["ctr"] = ctr

	return stats, nil
}
