package seeds

import (
	"log"
	"time"

	"github.com/bagus123/go-rest-mongodb/daos"
	"github.com/bagus123/go-rest-mongodb/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// RunUserSeed ...
func RunUserSeed() {
	newUser := models.User{
		ID:        primitive.NewObjectID(),
		UserName:  "john99",
		Password:  "123",
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, _ := daos.UserDAO().InsertOne(newUser)
	log.Println(result)

	newUser = models.User{
		ID:        primitive.NewObjectID(),
		UserName:  "alien12",
		Password:  "123",
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, _ = daos.UserDAO().InsertOne(newUser)
	log.Println(result)

	newUser = models.User{
		ID:        primitive.NewObjectID(),
		UserName:  "spider8",
		Password:  "123",
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, _ = daos.UserDAO().InsertOne(newUser)
	log.Println(result)

	newUser = models.User{
		ID:        primitive.NewObjectID(),
		UserName:  "lion5",
		Password:  "123",
		Role:      "Admin",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	result, _ = daos.UserDAO().InsertOne(newUser)
	log.Println(result)

}
