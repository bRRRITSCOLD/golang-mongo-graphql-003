package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"golang-mongo-graphql-003/internal/api/generated"
	"golang-mongo-graphql-003/internal/data_management"
	"golang-mongo-graphql-003/internal/middleware"
	"golang-mongo-graphql-003/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *agentResolver) Authors(ctx context.Context, obj *models.Agent) ([]*models.Author, error) {
	// panic(fmt.Errorf("not implemented"))
	gc, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	return r.DataLoaders.Retrieve(gc.Request.Context()).AuthorsByAgentIds.Load(obj.AgentID)
}

func (r *mutationResolver) CreateAgent(ctx context.Context, input models.CreateAgentInput) (*models.Agent, error) {
	createdAgents, createAgentsErr := data_management.CreateAgents([]models.Agent{
		{
			Name:  input.Name,
			Email: input.Email,
			// AuthorIDs: input.AuthorIDs,
		},
	})
	if createAgentsErr != nil {
		return nil, createAgentsErr
	}
	return models.PointerAgent(createdAgents[0]), nil
}

func (r *queryResolver) Agent(ctx context.Context, agentID string) (*models.Agent, error) {
	foundAgents, findAgentsErr := data_management.FindAgents(models.Agent{
		AgentID: agentID,
	})
	if findAgentsErr != nil {
		return nil, findAgentsErr
	}
	return models.PointerAgent(foundAgents[0]), nil
}

func (r *queryResolver) Agents(ctx context.Context) ([]*models.Agent, error) {
	foundAgents, findAgentsErr := data_management.FindAgents(bson.D{})
	if findAgentsErr != nil {
		return nil, findAgentsErr
	}
	return models.PointerAgents(foundAgents), nil
}

// Agent returns generated.AgentResolver implementation.
func (r *Resolver) Agent() generated.AgentResolver { return &agentResolver{r} }

type agentResolver struct{ *Resolver }
