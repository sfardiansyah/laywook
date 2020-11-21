package mongodb

import (
	"context"
	"os"
	"time"

	"github.com/sfardiansyah/laywook/pkg/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage ...
type Storage struct {
	db *mongo.Database
}

// NewStorage ...
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	c, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGODB_URL")))
	if err != nil {
		return nil, err
	}

	s.db = c.Database(os.Getenv("MONGODB_DB"))

	return s, nil
}

// GetUser ...
func (s *Storage) GetUser(email string) (*auth.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := new(auth.User)

	collection := s.db.Collection("users")

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, auth.ErrInvalidCredentials
	}

	return user, nil
}
