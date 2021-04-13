package dataloaders

//go:generate go run github.com/vektah/dataloaden AgentsByAgentIdsLoader string *golang-mongo-graphql-003/internal/models.Agent

import (
	"context"
	"golang-mongo-graphql-003/internal/data_management"
	"golang-mongo-graphql-003/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type Loaders struct {
	// individual loaders will be defined here
	AgentsByAgentIds *AgentsByAgentIdsLoader
}

func newLoaders(ctx context.Context) *Loaders {
	return &Loaders{
		AgentsByAgentIds: newAgentsByAgentIds(ctx),
	}
}

// Retriever retrieves dataloaders from the request context.
type Retriever interface {
	Retrieve(context.Context) *Loaders
}

type retriever struct {
	key contextKey
}

func (r *retriever) Retrieve(ctx context.Context) *Loaders {
	return ctx.Value(r.key).(*Loaders)
}

// NewRetriever instantiates a new implementation of Retriever.
func NewRetriever() Retriever {
	return &retriever{key: key}
}

func newAgentsByAgentIds(ctx context.Context) *AgentsByAgentIdsLoader {
	return NewAgentsByAgentIdsLoader(AgentsByAgentIdsLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(agentIDs []string) ([]*models.Agent, []error) {
			agents, err := data_management.FindAgents(bson.D{
				bson.E{"agentId", bson.D{
					bson.E{"$in", append(bson.A{}, agentIDs)},
				}},
			})
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByAuthorID := make(map[string]*models.Agent, len(agentIDs))
			for _, agent := range agents {
				groupByAuthorID[agent.AgentID] = models.PointerAgent(agent)
			}
			// order
			result := make([]*models.Agent, len(agentIDs))
			for i, agentID := range agentIDs {
				result[i] = groupByAuthorID[agentID]
			}
			return result, nil
		},
	})
}
