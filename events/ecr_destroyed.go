package events

// ECR destroyed event
type ECRDestroyedEvent struct {
	GeneralEvent
	Payload ECRDestroyedPayload `json:"payload"`
}

// ECR destroyed event payload
type ECRDestroyedPayload struct {
	Name       string            `json:"name"`
	Kind       string            `json:"kind"`
	ApiVersion string            `json:"apiVersion"`
	Labels     map[string]string `json:"labels"`
	Properties map[string]string `json:"properties"`
}
