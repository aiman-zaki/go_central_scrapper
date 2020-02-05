package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//Timetable struct
type Timetable struct {
	Group  string
	Start  string
	End    string
	Day    string
	Mode   string
	Status string
	Room   string
}

//Course Represents
type Course struct {
	Code       string      `json:"code"`
	Timetables []Timetable `json:"timetables"`
}

//Faculty represents
type Faculty struct {
	Code    string   `json:"code"`
	Name    string   `json:"name"`
	Courses []Course `json:"courses"`
}

//ReturnOneFaculty is quury a faculty
func ReadOneFaculty(client *mongo.Client, filter bson.M) Faculty {
	var faculty Faculty
	collection := client.Database("scrapper").Collection("icress")
	documentReturned := collection.FindOne(context.TODO(), filter)
	documentReturned.Decode(&faculty)
	return faculty
}

//InsertNewFaculty is insert new faculty
func InsertNewFaculty(client *mongo.Client, faculty Faculty) interface{} {
	collection := client.Database("scrapper").Collection("icress")
	insertResult, err := collection.InsertOne(context.TODO(), faculty)
	if err != nil {
		log.Fatalln("Error on inserting new Hero", err)
	}
	return insertResult.InsertedID
}
