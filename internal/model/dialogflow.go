package model

// DialogFlowRequest is the incoming body that is sent to the webhook's POST route.
type DialogFlowRequest struct {
	ResponseID  string `json:"responseId,omitempty"`
	QueryResult struct {
		QueryText  string `json:"queryText,omitempty"`
		Action     string `json:"action,omitempty"`
		Parameters struct {
			Name string `json:"name,omitempty"`
		} `json:"parameters,omitempty"`
		AllRequiredParamsPresent bool   `json:"allRequiredParamsPresent,omitempty"`
		FulfillmentText          string `json:"fulfillmentText,omitempty"`
		FulfillmentMessages      []struct {
			Text struct {
				Text []string `json:"text,omitempty"`
			} `json:"text,omitempty"`
		} `json:"fulfillmentMessages,omitempty"`
		OutputContexts []struct {
			Name          string `json:"name,omitempty"`
			LifespanCount int    `json:"lifespanCount,omitempty"`
			Parameters    struct {
				Name         string `json:"name,omitempty"`
				NameOriginal string `json:"name.original,omitempty"`
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
			GroupID       string `json:"groupId,omitempty"`
			BotID         string `json:"botId,omitempty"`
			Attachments   []any  `json:"attachments,omitempty"`
			KmUserLocale  string `json:"kmUserLocale,omitempty"`
			From          string `json:"from,omitempty"`
			ApplicationID string `json:"applicationId,omitempty"`
			MessageSource string `json:"messageSource,omitempty"`
		} `json:"payload,omitempty"`
	} `json:"originalDetectIntentRequest,omitempty"`
	Session string `json:"session,omitempty"`
}

// A rich response card which contains a button, an image, and text.
type RichResponseCard struct {
	FulfillmentMessages []FulfillmentMessageCard `json:"fulfillmentMessages,omitempty"`
}

type RichResponseButton struct {
	FulfillmentMessages []Button `json:"fulfillmentMessages,omitempty"`
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

type FulfillmentMessageCard struct {
	Card Card `json:"card,omitempty"`
}

type Parameters struct {
	Name string `json:"name,omitempty"`
}
