package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"logger/api"
	"logger/model"
	"net/http"
	"os"
)

const (
	webPort  = "80"
	rpcPort  = "5001"
	gRpcPort = "50001"
)

// "mongodb://mongo:27017"
func connectToMongo() (*mongo.Client, error) {
	// create connection options
	mongoURL := os.Getenv("mongoURL")
	dbUser := os.Getenv("dbUser")
	dbPassword := os.Getenv("dbPassword")

	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: dbUser,
		Password: dbPassword,
	})

	// connect
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	//// Ping the MongoDB server to verify the connection
	//err = client.Ping(context.Background(), nil)
	//if err != nil {
	//	log.Fatal(err)
	//}

	log.Println("Connected to mongo!")

	return client, nil
}

func onDBConnected(s *mongo.Database) error {
	fmt.Println("Connected to DB " + s.Name())

	model.InitLogEntryModel(s)

	return nil
}

func main() {
	client, err := connectToMongo()
	if err != nil {
		return
	}
	dbName := os.Getenv("dbName")

	err = onDBConnected(client.Database(dbName))
	if err != nil {
		return
	}

	go api.GRPCListen()

	protocol := os.Getenv("protocol")
	if protocol == "" {
		//protocol = "THRIFT"
		protocol = "HTTP"
	}

	if protocol == "HTTP" {
		srv := &http.Server{
			Addr:    ":8089",
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
