package repositories

import (
	"context"
	"errors"
	domain "task_management/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// define userRepsoitory interface
type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	FindByUsername(ctx context.Context, username string) (*domain.User, error)
	CountByUsername(ctx context.Context, username string) (int64, error)
	CountAll(ctx context.Context) (int64, error)
	PromoteUser(ctx context.Context, userID string) error
}

// implementation of userrepository and uses monogb
type UserRespositoryImpl struct {
	Collection *mongo.Collection
}


// constructor to initialize userRepositoryImpl
func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &UserRespositoryImpl{Collection: collection}
}

// inserts a new user to mongodb collection
func (r *UserRespositoryImpl) CreateUser(ctx context.Context, user *domain.User) error {
	_, err := r.Collection.InsertOne(ctx, user)
	return err
}

// retrieves a user based on the given username
func (r *UserRespositoryImpl) FindByUsername(ctx context.Context, username string) (*domain.User, error) {
	
	var user domain.User
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}


// counts the number of users that matches the username
func (r *UserRespositoryImpl) CountByUsername(ctx context.Context, username string) (int64, error) {
	return r.Collection.CountDocuments(ctx, bson.M{"username": username})
}

//counts the total number of user documents in the collection

func (r *UserRespositoryImpl) CountAll(ctx context.Context) (int64, error) {
	return r.Collection.CountDocuments(ctx, bson.M{})
}

// updates the user role to admin based the id provided
func (r *UserRespositoryImpl) PromoteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user id")
	}
	//filters the id
	filter := bson.M{"_id": objID}
	// the thing to be updated
	update := bson.M{"$set": bson.M{"role": "Admin"}}

	result, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	//if no result or the count is 0 it returns error
	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}
	return nil
}
