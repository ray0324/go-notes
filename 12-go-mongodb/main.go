package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Person 用户信息
type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

// CreatePersonEndpoint 创建一个用户
func CreatePersonEndpoint(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("content-type", "application/json")
	var person Person

	json.NewDecoder(req.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	rst, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(res).Encode(rst)
}

func main() {
	fmt.Println("Starting the app ...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	router := mux.NewRouter()

	router.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	http.ListenAndServe(":8080", router)

}
