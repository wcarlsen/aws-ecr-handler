package events

// ECR added event
type ECRAddedEvent struct {
	GeneralEvent
	Payload ECRAddedPayload `json:"payload"`
}

// ECR added event payload
type ECRAddedPayload struct {
	Name       string                    `json:"name"`
	Kind       string                    `json:"kind"`
	ApiVersion string                    `json:"apiVersion"`
	Properties ECRAddedPayloadProperties `json:"properties"`
}

type ECRAddedPayloadProperties struct {
	RepoName  string `json:"repoName"`
	AccountId string `json:"accountId"`
}
