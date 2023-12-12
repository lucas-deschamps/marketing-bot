package chatbot

import (
	"fmt"

	"github.com/lucas-deschamps/marketing-bot/internal/model"
)

func CouponOffer(parameters model.Parameters) model.SuggestionButton {
	customerName := parameters.Name

	msg := fmt.Sprintf("It's a pleasure to meet you, %s. Are you interested in our coupon promotion?", customerName)

	couponMessage := model.SuggestionButton{
		Metadata: model.Metadata{
			ContentType: "300",
			TemplateID:  "6",
			Payload: []model.Payload{
				{
					Message: "Yes, I'm interested!",
					Title:   "Yes! Show me the coupon",
				},
				{
					Message: "No, thanks, I'm not interested.",
					Title:   "No, thanks",
				},
			},
		},
		Message:  msg,
		Platform: "kommunicate",
	}

	return couponMessage
}

func CouponYes(parameters model.Parameters) model.RichResponseCard {
	couponLinkMsg := model.RichResponseCard{
		FulfillmentMessages: []model.FulfillmentMessageCard{
			{
				Card: model.Card{
					Title:    "Here is our unique promotional coupon!",
					Subtitle: " 10% off. Limit 1 per customer.",
					ImageURI: "",
					Buttons: []model.Button{
						{
							Text:     "Reveal Coupon",
							Postback: "https://9415-2804-14d-5c5b-8a57-3a53-7f8d-8f07-4718.ngrok-free.app/v1/coupon",
						},
					},
				},
			},
		},
	}

	return couponLinkMsg
}

func CouponNo(parameters model.Parameters) model.RichResponseCard {
	goodbyeMsg := fmt.Sprintf("No worries! Have a nice day!")

	couponRefusalCard := model.RichResponseCard{
		FulfillmentMessages: []model.FulfillmentMessageCard{
			{
				Card: model.Card{
					Title:    "See you!",
					Subtitle: goodbyeMsg,
					ImageURI: "https://i.imgur.com/82p9aPV.jpeg",
					Buttons:  []model.Button{},
				},
			},
		},
	}

	return couponRefusalCard
}
