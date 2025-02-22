package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/radhika.parmar/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://radhikaparmar:U5WkCeAeaRt4h6uT@cluster0.v0f60.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbName = "netflix"
const colName = "watchlist"

// Set connection
var collection *mongo.Collection // It is the reference of table on mongodb

// connect with mongodb
func init() {
	clientOpetions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOpetions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection sucess..")

	collection = client.Database(dbName).Collection(colName)

	fmt.Println("Collection reference is ready")
}

// mongodb helpers
// insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

// update one record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete one record
func deleteOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", result.DeletedCount)
}

// delete all records from mongodb
func deleteAllMovies() {
	// collection.DeleteMany(context.Background(),bson.D{{}})
	result, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted count: ", result.DeletedCount)
}

// get all movies
// func getAllMovies() {	
// }
