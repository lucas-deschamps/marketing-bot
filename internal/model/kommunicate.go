package model

// Docs: https://docs.kommunicate.io/docs/message-types
type SuggestionButton struct {
	Message  string   `json:"message,omitempty"`
	Platform string   `json:"platform,omitempty"`
	Metadata Metadata `json:"metadata,omitempty"`
}

type Payload struct {
	Title   string `json:"title,omitempty"`
	Message string `json:"message,omitempty"`
}

type Metadata struct {
	ContentType string    `json:"contentType,omitempty"`
	TemplateID  string    `json:"templateId,omitempty"`
	Payload     []Payload `json:"payload,omitempty"`
}

type LinkButton struct {
	Message  string       `json:"message,omitempty"`
	Platform string       `json:"platform,omitempty"`
	Metadata LinkMetadata `json:"metadata,omitempty"`
}

type LinkPayload struct {
	Type             string `json:"type,omitempty"`
	URL              string `json:"url,omitempty"`
	Name             string `json:"name,omitempty"`
	OpenLinkInNewTab bool   `json:"openLinkInNewTab,omitempty"`
}

type LinkMetadata struct {
	ContentType string        `json:"contentType,omitempty"`
	TemplateID  string        `json:"templateId,omitempty"`
	Payload     []LinkPayload `json:"payload,omitempty"`
}

type SubmitButton struct {
	Message  string         `json:"message,omitempty"`
	Platform string         `json:"platform,omitempty"`
	Metadata SubmitMetadata `json:"metadata,omitempty"`
}

type SubmitPayload struct {
	Name      string `json:"name,omitempty"`
	ReplyText string `json:"replyText,omitempty"`
}

type FormData struct {
	Amount      string `json:"amount,omitempty"`
	Description string `json:"description,omitempty"`
}

type SubmitMetadata struct {
	ContentType string          `json:"contentType,omitempty"`
	TemplateID  string          `json:"templateId,omitempty"`
	Payload     []SubmitPayload `json:"payload,omitempty"`
	FormData    FormData        `json:"formData,omitempty"`
	FormAction  string          `json:"formAction,omitempty"`
	RequestType string          `json:"requestType,omitempty"`
}
