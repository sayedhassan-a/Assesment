package utils

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

type dbConfig struct {
	URI string `yaml:"uri"`
	DBName string `yaml:"dbName"`
	Collection []string `yaml:"collection"`
}

func LoadDBConfig() dbConfig{

	data, err := ioutil.ReadFile("./config/database-config.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)

	}

	var config dbConfig
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return config

}
func ConnectToDB() *mongo.Client{

	ctx := context.TODO();

	config := LoadDBConfig()
	option := options.Client().ApplyURI(config.URI)
	mongoClient, err := mongo.Connect(ctx,option)
	if err != nil {
		log.Fatalf("error connecting to the database: %v", err)
	}
	println(mongoClient)
	println("connected")
	return mongoClient

}
var DB = ConnectToDB()
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("appdb").Collection(collectionName)
	return collection
}