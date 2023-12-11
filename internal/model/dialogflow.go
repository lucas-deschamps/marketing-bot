package model

// DialogFlowRequest is the incoming body that is sent to the webhook's POST route.
type DialogFlowRequest struct {
	ResponseID  string `json:"responseId,omitempty"`
	QueryResult struct {
		QueryText  string `json:"queryText,omitempty"`
		Action     string `json:"action,omitempty"`
		Parameters struct {
			Email string `json:"email,omitempty"`
		} `json:"parameters,omitempty"`
		AllRequiredParamsPresent bool `json:"allRequiredParamsPresent,omitempty"`
		OutputContexts           []struct {
			Name       string `json:"name,omitempty"`
			Parameters struct {
				NoInput       float64 `json:"no-input,omitempty"`
				NoMatch       float64 `json:"no-match,omitempty"`
				Email         string  `json:"email,omitempty"`
				EmailOriginal string  `json:"email.original,omitempty"`
			} `json:"parameters,omitempty"`
		} `json:"outputContexts,omitempty"`
		Intent struct {
			Name        string `json:"name,omitempty"`
			DisplayName string `json:"displayName,omitempty"`
		} `json:"intent,omitempty"`
		IntentDetectionConfidence float64 `json:"intentDetectionConfidence,omitempty"`
		LanguageCode              string  `json:"languageCode,omitempty"`
	} `json:"queryResult,omitempty"`
	OriginalDetectIntentRequest struct {
		Payload struct {
		} `json:"payload,omitempty"`
	} `json:"originalDetectIntentRequest,omitempty"`
	Session string `json:"session,omitempty"`
}

type RichResponse struct {
	FulfillmentMessages []FulfillmentMessages `json:"fulfillmentMessages,omitempty"`
}

type Button struct {
	Text     string `json:"text,omitempty"`
	Postback string `json:"postback,omitempty"`
}

type Card struct {
	Title    string   `json:"title,omitempty"`
	Subtitle string   `json:"subtitle,omitempty"`
	ImageURI string   `json:"imageUri,omitempty"`
	Buttons  []Button `json:"buttons,omitempty"`
}

type FulfillmentMessages struct {
	Card Card `json:"card,omitempty"`
}
