package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"unicode"

	scrapper "github.com/aiman-zaki/go_central_scrapper/scrappers"

	models "github.com/aiman-zaki/go_central_scrapper/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://admin:password@localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func removeSpace(s string) string {
	rr := make([]rune, 0, len(s))
	for _, r := range s {
		if !unicode.IsSpace(r) {
			rr = append(rr, r)
		}
	}
	return string(rr)
}

func main() {

	//MongoDb Connection
	c := getClient()
	err := c.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("Couldn't connect to the database", err)
	} else {
		log.Println("Connected!")
	}

	faculties := scrapper.FetchFaculties()
	for index, faculty := range faculties {
		temp := models.ReadOneFaculty(c, bson.M{"code": faculty.Code})
		if (reflect.DeepEqual(temp, models.Faculty{})) {
			faculties[index] = scrapper.FetchCourses(faculty)
			insertedID := models.InsertNewFaculty(c, faculties[index])
			log.Println(insertedID)
		} else {
			log.Println("Entry exist in Database")
		}

	}
	fmt.Printf("%+v\n", faculties)

}
