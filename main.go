package main

import (
	"LOGGER-SERVICE/api"
	"LOGGER-SERVICE/model"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
)

// "mongodb://mongo:27017"
func connectToMongo() (*mongo.Database, error) {
	// create connection options
	mongoURL := os.Getenv("mongoURL")
	admin := os.Getenv("mongoURL")
	password := os.Getenv("mongoURL")
	dbName := os.Getenv("dbName")

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: admin,
		Password: password,
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo!")

	return c.Database(dbName), nil
}

func onDBConnected(s *mongo.Database) error {
	fmt.Println("Connected to DB " + s.Name())

	model.InitLogEntryModel(s)

	return nil
}

func main() {
	db, err := connectToMongo()
	if err != nil {
		return
	}

	err = onDBConnected(db)
	if err != nil {
		return
	}

	protocol := os.Getenv("protocol")
	if protocol == "" {
		//protocol = "THRIFT"
		protocol = "HTTP"
	}

	if protocol == "HTTP" {
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%s", webPort),
			Handler: api.Routes(),
		}

		err := srv.ListenAndServe()
		if err != nil {
			log.Panic(err)
		}
	} else {
		return
	}
}
