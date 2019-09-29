package daos

import (
	"context"
	"log"
	"time"

	"github.com/bagus123/go-rest-mongodb/models"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserDAOImpl ...
type UserDAOImpl struct {
	IDAO
}

var db *mongo.Database
var col *mongo.Collection

// TableName ...
const TableName = "users"

// Init ...
func (m *UserDAOImpl) Init(database *mongo.Database) {
	db = database
	col = db.Collection(TableName)
}

// User ...
type User = models.User

// GetAll ...
func (m *UserDAOImpl) GetAll() ([]interface{}, error) {
	var users []interface{}
	filter := bson.M{}
	result, err := col.Find(context.TODO(), filter)
	if err != nil {
		return nil, errors.Wrap(err, "Error Find User")
	}

	// Iterate through the returned cursor.
	for result.Next(context.TODO()) {
		var user User
		err = result.Decode(&user)
		if err != nil {
			return nil, errors.Wrap(err, "Error Decode document")
		}
		log.Println(user)
		users = append(users, user)
	}

	log.Println(users)
	return users, nil
}

// InsertOne ...
func (m *UserDAOImpl) InsertOne(user interface{}) (interface{}, error) {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := col.InsertOne(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "Error InsertOne")
	}
	return result.InsertedID, nil
}

// GetOneByID ...
func (m *UserDAOImpl) GetOneByID(id string) (interface{}, error) {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	result := col.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		return nil, errors.Wrap(err, "Error GetOneByID")
	}

	var user User
	err = result.Decode(&user)
	if err != nil {
		return nil, errors.Wrap(err, "Error GetOneByID")
	}

	return user, nil
}

// UpdateOne ...
func (m *UserDAOImpl) UpdateOne(filter bson.M, user interface{}) (interface{}, error) {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updateData := bson.M{"$set": user}
	result, err := col.UpdateOne(ctx, filter, updateData)

	if err != nil {
		return nil, errors.Wrap(err, "Error UpdateOne")
	}

	return result.ModifiedCount, nil
}

// UpdateOneByID ...
func (m *UserDAOImpl) UpdateOneByID(id string, user interface{}) (interface{}, error) {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	updateData := bson.M{"$set": user}
	result, err := col.UpdateOne(ctx, filter, updateData)

	if err != nil {
		return nil, errors.Wrap(err, "Error UpdateOneByID")
	}

	return result.ModifiedCount, nil
}

// DeleteOneByID ...
func (m *UserDAOImpl) DeleteOneByID(id string) (interface{}, error) {

	// create a new context with a 10 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	result, err := col.DeleteOne(ctx, filter)

	if err != nil {
		return nil, errors.Wrap(err, "Error DeleteOneByID")
	}

	return result.DeletedCount, nil
}
