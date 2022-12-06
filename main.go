package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"
)

type Post struct {
	Title string `bson:"title,omitempty"`
	Body  string `bson:"body,omitempty"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:example@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second)
	client.Connect(ctx)
	defer client.Disconnect(ctx)
	collection := client.Database("blog").Collection("posts")
	docs := []interface{}{
		bson.D{{"title", "Danik"}, {"world", "Suzdaleu"}},
		bson.D{{"title", "Kirill"}, {"world", "Kaluga"}},
	}
	res, err := collection.InsertMany(ctx, docs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
