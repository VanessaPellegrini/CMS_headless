package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/VanessaPellegrini/CMS_headless/app/shared/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ()

func GetCollection(collection string) *mongo.Collection {
	get := utils.GetEnvWithKey
	USR := get("MONGODB_USERNAME")
	PWD := get("MONGODB_PW")
	HOST := get("MONGODB_HOST")
	PORT := get("MONGODB_PORT")
	DBNAME := get("MONGODB_NAME")

	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s", USR, PWD, HOST, PORT)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	log.Printf("uri %s", uri)

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(DBNAME).Collection(collection)
}
