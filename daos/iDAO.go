package daos

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// IDAO ...
type IDAO interface {
	Init(database *mongo.Database)
	GetAll() ([]interface{}, error)
	GetOneByID(id string) (interface{}, error)
	InsertOne(interface{}) (interface{}, error)
	UpdateOne(filter bson.M, user interface{}) (interface{}, error)
	UpdateOneByID(id string, user interface{}) (interface{}, error)
	DeleteOneByID(id string) (interface{}, error)
}
