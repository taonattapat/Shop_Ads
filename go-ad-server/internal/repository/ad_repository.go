package repository

import (
	"context"
	"time"

	"go-ad-server/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdRepository struct {
	Collection *mongo.Collection
}

func NewAdRepository(db *mongo.Database) *AdRepository {
	return &AdRepository{
		Collection: db.Collection("ads"),
	}
}

func (r *AdRepository) Create(ad *models.Ad) error {
	ad.CreatedAt = time.Now()
	ad.UpdatedAt = time.Now()
	ad.Status = "active" // Default status
	ad.Spent = 0

	result, err := r.Collection.InsertOne(context.Background(), ad)
	if err != nil {
		return err
	}
	ad.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (r *AdRepository) GetAds(filter bson.M) ([]models.Ad, error) {
	var ads []models.Ad
	cursor, err := r.Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.Background(), &ads); err != nil {
		return nil, err
	}
	return ads, nil
}

func (r *AdRepository) GetByID(id string) (*models.Ad, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var ad models.Ad
	err = r.Collection.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&ad)
	if err != nil {
		return nil, err
	}
	return &ad, nil
}

func (r *AdRepository) UpdateStatus(id string, status string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	update := bson.M{
		"$set": bson.M{
			"status":     status,
			"updated_at": time.Now(),
		},
	}
	_, err = r.Collection.UpdateOne(context.Background(), bson.M{"_id": objID}, update)
	return err
}

func (r *AdRepository) IncrementSpent(id primitive.ObjectID, amount float64) error {
	update := bson.M{
		"$inc": bson.M{"spent": amount},
	}
	_, err := r.Collection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	return err
}

// For active ads, we might want a specialized query
func (r *AdRepository) GetActiveAds() ([]models.Ad, error) {
	filter := bson.M{
		"status": "active",
		"$expr":  bson.M{"$lt": []string{"$spent", "$budget"}}, // spent < budget
	}
	return r.GetAds(filter)
}
