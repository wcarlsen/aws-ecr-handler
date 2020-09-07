package events

// ECR provisioned event
type ECRProvisionedEvent struct {
	GeneralEvent
	Payload ECRProvisionedPayload `json:"payload"`
}

// ECR provisioned event payload
type ECRProvisionedPayload struct {
	Name       string            `json:"name"`
	Kind       string            `json:"kind"`
	ApiVersion string            `json:"apiVersion"`
	Labels     map[string]string `json:"labels"`
	Properties map[string]string `json:"properties"`
}
