package events

// General event
type GeneralEvent struct {
	Id            string      `json:"id"`
	CorrelationId string      `json:"correlationId"`
	Type          string      `json:"type"`
	SchemaVersion string      `json:"schemaVersion"`
	CreationTime  string      `json:"x-creationTime"`
	PayloadSchema string      `json:"x-payload-schema"`
	Payload       interface{} `json:"payload"`
}
