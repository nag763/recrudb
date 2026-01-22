package offer

import "nag763/recrudb/api/internal/domain"

type Offer struct {
	EnterpriseName        string         `json:"enterpriseName"`
	Contact               domain.Contact `json:"contact"`
	Address               domain.Address `json:"address"`
	TechnicalRequirements []string       `json:"techRequirements"`
	CandidateScore        float32        `json:"candidateScore"`
	HowToApply            string         `json:"howToApply"`
	Description           string         `json:"description"`
	AgentRecap            string         `json:"agentRecap"`
}
