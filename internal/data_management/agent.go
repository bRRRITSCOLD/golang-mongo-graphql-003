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

//CreateAgents - Insert multiple documents at once in the collection.
func CreateAgents(agents []models.Agent) ([]models.Agent, error) {
	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var newAgents []models.Agent
	for _, v := range agents {
		newAgent := models.NewAgent(v)
		timeNow := time.Now()
		newAgent.AgentID = uuid.New().String()
		newAgent.CreatedDate = timeNow
		newAgent.UpdatedDate = timeNow
		newAgents = append(newAgents, newAgent)
	}

	//Map struct slice to interface slice as InsertMany accepts interface slice as parameter
	var insertableAgents []interface{}
	// insertableAgents := make([]interface{}, len(newAgents))
	for _, v := range newAgents {
		insertableAgents = append(insertableAgents, v)
	}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return newAgents, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.AGENTS)

	//Perform InsertMany operation & validate against the error.
	insertManyResponse, err := collection.InsertMany(context.TODO(), insertableAgents)
	if err != nil {
		return newAgents, err
	}

	// create items to be returned
	var returnableAgents []models.Agent
	for i := 0; i < len(insertManyResponse.InsertedIDs); i++ {
		newAgent := models.NewAgent(newAgents[i])
		newAgentID := insertManyResponse.InsertedIDs[i]
		newAgent.ID = newAgentID.(primitive.ObjectID).Hex()
		returnableAgents = append(returnableAgents, newAgent)
	}

	//Return success without any error.
	return returnableAgents, nil
}

//FindAgents - Get All agents that match a criteria for collection
func FindAgents(filter interface{}) ([]models.Agent, error) {
	agents := []models.Agent{}

	//Get MongoDB connection using connectionhelper.
	client, err := mongodb.GetMongoClient()
	if err != nil {
		return agents, err
	}

	//Create a handle to the respective collection in the database.
	collection := client.Database(mongodb.DB).Collection(mongodb.AGENTS)

	//Perform Find operation & validate against the error.
	cur, findError := collection.Find(context.TODO(), filter)
	if findError != nil {
		return agents, findError
	}

	//Map result to slice
	for cur.Next(context.TODO()) {
		t := models.Agent{}

		err := cur.Decode(&t)
		if err != nil {
			return agents, err
		}

		agents = append(agents, t)
	}

	// once exhausted, close the cursor
	cur.Close(context.TODO())

	if len(agents) == 0 {
		return agents, mongo.ErrNoDocuments
	}
	return agents, nil
}
