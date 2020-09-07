package events

// ECR deleted event
type ECRDeletedEvent struct {
	GeneralEvent
	Payload ECRDeletedPayload `json:"payload"`
}

// ECR deleted event payload
type ECRDeletedPayload struct {
	Name       string `json:"name"`
	Kind       string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	// Labels     map[string]string         `json:"labels"`
	Properties ECRAddedPayloadProperties `json:"properties"`
}

// ECR deleted event properites
type ECRDeletedPayloadProperties struct {
	RepoName  string `json:"repoName"`
	AccountId string `json:"accountId"`
}
