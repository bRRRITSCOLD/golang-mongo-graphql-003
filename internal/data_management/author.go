package data_management

import (
	"context"
	"time"

	"golang-mongo-graphql-003/internal/models"
	"golang-mongo-graphql-003/internal/mongodb"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//CreateAuthors - Insert multiple documents at once in the collection.
func CreateAuthors(authors []models.Author) ([]models.Author, error) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var newAuthors []models.Author
	for _, v := range authors {
		newAuthor := models.NewAuthor(v)
		timeNow := time.Now()
		newAuthor.AuthorID = uuid.New().String()
		newAuthor.CreatedDate = timeNow
		newAuthor.UpdatedDate = timeNow
		newAuthors = append(newAuthors, newAuthor)
	}

	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var insertableAuthors []interface{}
	// insertableAuthors := make([]interface{}, len(newAuthors))
	for _, v := range newAuthors {
		insertableAuthors = append(insertableAuthors, v)
	}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return newAuthors, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.AUTHORS)

	//Perform InsertMany operation & validate against the error.
	insertManyResponse, err := collection.InsertMany(context.TODO(), insertableAuthors)
	if err != nil {
		return newAuthors, err
	}

	// create items to be returned
	var returnableAuthors []models.Author
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newAuthor := models.NewAuthor(newAuthors[i])
		newAuthorID := insertManyResponse.InsertedIDs[i]
		newAuthor.ID = newAuthorID.(primitive.ObjectID).Hex()
		returnableAuthors = append(returnableAuthors, newAuthor)
	}

	//Return success without any error.
	return returnableAuthors, nil
}

//FindAuthors - Get All authors that match a criteria for collection
func FindAuthors(filter interface{}) ([]models.Author, error) {
	authors := []models.Author{}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return authors, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.AUTHORS)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return authors, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := models.Author{}

		err := cur.Decode(&t)
		if err != nil {
			return authors, err
		}

		authors = append(authors, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(authors) == 0 {
		return authors, mongo.ErrNoDocuments
	}
	return authors, nil
}
