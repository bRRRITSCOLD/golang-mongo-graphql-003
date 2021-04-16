package dataloaders

// go:generate go run github.com/vektah/dataloaden AgentsByAuthorIdsLoader string *golang-mongo-graphql-003/internal/models.Agent
// go:generate go run github.com/vektah/dataloaden AuthorsByAgentIdsLoader string []*golang-mongo-graphql-003/internal/models.Author
import (
	"context"
	"golang-mongo-graphql-003/internal/data_management"
	"golang-mongo-graphql-003/internal/models"
	"golang-mongo-graphql-003/internal/utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type contextKey string

const key = contextKey("dataloaders")

// Loaders holds references to the individual dataloaders.
type Loaders struct {
	// individual loaders will be defined here
	AgentsByAuthorIds *AgentsByAuthorIdsLoader
	AuthorsByAgentIds *AuthorsByAgentIdsLoader
}

func newLoaders(ctx context.Context) *Loaders {
	return &Loaders{
		AgentsByAuthorIds: newAgentsByAuthorIds(ctx),
		AuthorsByAgentIds: newAuthorsByAgentIds(ctx),
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

func newAgentsByAuthorIds(ctx context.Context) *AgentsByAuthorIdsLoader {
	return NewAgentsByAuthorIdsLoader(AgentsByAuthorIdsLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(authorIDs []string) ([]*models.Agent, []error) {
			agents, err := data_management.FindAgents(bson.D{
				bson.E{"authorIds", bson.D{
					bson.E{"$in", authorIDs},
				}},
			})
			if err != nil {
				return nil, []error{err}
			}
			// map
			groupByAgentID := make(map[string]*models.Agent, len(authorIDs))
			for _, agent := range agents {
				for _, authorID := range authorIDs {
					if utils.ContainsString(agent.AuthorIDs, authorID) {
						groupByAgentID[authorID] = models.PointerAgent(agent)
					}
				}
			}
			// order
			result := make([]*models.Agent, len(authorIDs))
			for i, authorID := range authorIDs {
				result[i] = groupByAgentID[authorID]
			}
			return result, nil
		},
	})
}

func newAuthorsByAgentIds(ctx context.Context) *AuthorsByAgentIdsLoader {
	return NewAuthorsByAgentIdsLoader(AuthorsByAgentIdsLoaderConfig{
		MaxBatch: 100,
		Wait:     5 * time.Millisecond,
		Fetch: func(agentIDs []string) ([][]*models.Author, []error) {
			agents, err := data_management.FindAuthors(bson.D{
				bson.E{"agentId", bson.D{
					bson.E{"$in", agentIDs},
				}},
			})
			if err != nil {
				return nil, []error{err}
			}
			// group
			groupByAgentID := make(map[string][]*models.Author, len(agentIDs))
			for _, agent := range agents {
				groupByAgentID[agent.AgentID] = append(groupByAgentID[agent.AgentID], models.PointerAuthor(agent))
			}
			// order
			result := make([][]*models.Author, len(agentIDs))
			for i, agentID := range agentIDs {
				result[i] = groupByAgentID[agentID]
			}
			return result, nil
		},
	})
}
