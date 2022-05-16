package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func get_viper(key string) string {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func GetCollection(collection string) *mongo.Collection {

	USR := get_viper("MONGODB_USERNAME")
	PWD := get_viper("MONGODB_PW")
	HOST := get_viper("MONGODB_HOST")
	PORT := get_viper("MONGODB_PORT")
	//DBNAME := viper.Get("MONGODB_NAME")

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

	return client.Database("docker").Collection(collection)
}
