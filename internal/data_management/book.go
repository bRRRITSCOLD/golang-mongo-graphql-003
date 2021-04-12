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

//CreateBooks - Insert multiple documents at once in the collection.
func CreateBooks(books []models.Book) ([]models.Book, error) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var newBooks []models.Book
	for _, v := range books {
		newBook := models.NewBook(v)
		timeNow := time.Now()
		newBook.BookID = uuid.New().String()
		newBook.CreatedDate = timeNow
		newBook.UpdatedDate = timeNow
		newBooks = append(newBooks, newBook)
	}

	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var insertableBooks []interface{}
	// insertableBooks := make([]interface{}, len(newBooks))
	for _, v := range newBooks {
		insertableBooks = append(insertableBooks, v)
	}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return newBooks, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.BOOKS)

	//Perform InsertMany operation & validate against the error.
	insertManyResponse, err := collection.InsertMany(context.TODO(), insertableBooks)
	if err != nil {
		return newBooks, err
	}

	// create items to be returned
	var returnableBooks []models.Book
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newBook := models.NewBook(newBooks[i])
		newBookID := insertManyResponse.InsertedIDs[i]
		newBook.ID = newBookID.(primitive.ObjectID).Hex()
		returnableBooks = append(returnableBooks, newBook)
	}

	//Return success without any error.
	return returnableBooks, nil
}

//FindBooks - Get All books that match a criteria for collection
func FindBooks(filter interface{}) ([]models.Book, error) {
	books := []models.Book{}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return books, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.BOOKS)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return books, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := models.Book{}

		err := cur.Decode(&t)
		if err != nil {
			return books, err
		}

		books = append(books, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(books) == 0 {
		return books, mongo.ErrNoDocuments
	}
	return books, nil
}
