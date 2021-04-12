package models

import (
	"time"
)

//Issue - struct to map with mongodb documents
type Agent struct {
	ID          string    `json:"_id" bson:"_id,omitempty"`
	AgentID     string    `json:"agentId" bson:"agentId,omitempty"`
	CreatedDate time.Time `json:"createdDate" bson:"createdDate,omitempty"`
	UpdatedDate time.Time `json:"updatedDate" bson:"updatedDate,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Email       string    `json:"email" bson:"email,omitempty"`
	AuthorIDs   []string  `json:"authorIds" bson:"authorIds,omitempty"`
}

type CreateAgentInput struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	AuthorIDs []string `json:"authorIds"`
}

func NewAgent(agent Agent) Agent {
	return Agent{
		ID:          agent.ID,
		AgentID:     agent.AgentID,
		CreatedDate: agent.CreatedDate,
		UpdatedDate: agent.UpdatedDate,
		Name:        agent.Name,
		Email:       agent.Email,
		AuthorIDs:   agent.AuthorIDs,
	}
}

func PointerAgent(agent Agent) *Agent {
	pointerAgent := agent
	return &pointerAgent
}

func PointerAgents(agents []Agent) []*Agent {
	var pointerAgents []*Agent
	for _, agent := range agents {
		pointerAgents = append(pointerAgents, PointerAgent(agent))
	}
	return pointerAgents
}
