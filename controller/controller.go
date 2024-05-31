package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skyakashh/mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbname          = "netflix"
	connectionsting = "mongodb+srv://sky:akash@cluster0.g6gxgfm.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	colname         = "watchlist"
)

var collection *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionsting)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection successful to MongoDB")

	collection = client.Database(dbname).Collection(colname)
	fmt.Println("Collection instance ready")
}

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted movie with ID:", inserted.InsertedID)
}

func updateOneRecord(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	res, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count:", res.ModifiedCount)
}

func deleteOneRecord(movieID string) {
	id, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with delete count:", res.DeletedCount)
}

func deleteAll() int {
	filter := bson.D{{}}
	res, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of deletions:", res.DeletedCount)
	return int(res.DeletedCount)
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(context.Background())

	var movies []primitive.M
	for cur.Next(context.TODO()) {
		var movie bson.M
		if err := cur.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	return movies
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "POST")
	var movie model.Netflix
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	updateOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneRecord(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	deleteCount := deleteAll()
	json.NewEncoder(w).Encode(deleteCount)
}
