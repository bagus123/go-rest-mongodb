package daos

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Base ..
type Base struct{}

func connect() *mongo.Database {
	options := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.NewClient(options)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	errPing := client.Ping(context.Background(), readpref.Primary())
	if errPing != nil {
		log.Fatal("Couldn't connect to the database", errPing)
	} else {
		log.Println("Connected!")
	}

	db := client.Database("exampleGoDB")
	return db
}

var mapDAOS map[string]IDAO

// InitDAO ...
func InitDAO() {

	db := connect()
	mapDAOS = make(map[string]IDAO)
	userDAO := new(UserDAOImpl)
	userDAO.Init(db)
	mapDAOS["UserDAO"] = userDAO

}

// UserDAO ...
func UserDAO() IDAO {
	return mapDAOS["UserDAO"]
}
