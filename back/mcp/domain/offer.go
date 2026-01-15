package domain

type Offer struct {
	EnterpriseName        string   `json:"enterpriseName"`
	Contact               Contact  `json:"contact"`
	Address               Address  `json:"address"`
	TechnicalRequirements []string `json:"techRequirements"`
	CandidateScore        float32  `json:"candidateScore"`
	HowToApply            string   `json:"howToApply"`
}
