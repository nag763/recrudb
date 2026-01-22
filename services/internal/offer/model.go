package offer

import (
	"nag763/recrudb/api/internal/address"
	"nag763/recrudb/api/internal/contact"
)

type Offer struct {
	EnterpriseName        string          `json:"enterpriseName"`
	Contact               contact.Contact `json:"contact"`
	Address               address.Address `json:"address"`
	TechnicalRequirements []string        `json:"techRequirements"`
	CandidateScore        float32         `json:"candidateScore"`
	HowToApply            string          `json:"howToApply"`
	Description           string          `json:"description"`
	AgentRecap            string          `json:"agentRecap"`
}
