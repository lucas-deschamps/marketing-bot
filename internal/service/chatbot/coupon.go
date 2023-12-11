package chatbot

import (
	"fmt"

	"github.com/lucas-deschamps/marketing-bot/internal/model"
)

func CouponService(parameters model.Parameters) model.SuggestionButton {
	customerName := parameters.Name

	msg := fmt.Sprintf("It's a pleasure to meet you, %s. Are you interested in our coupon promotion?", customerName)

	couponMessage := model.SuggestionButton{
		Metadata: model.Metadata{
			ContentType: "300",
			TemplateID:  "6",
			Payload: []model.Payload{
				{
					Message: "Yes, I'd like to!",
					Title:   "Yes! Show me the coupon",
				},
				{
					Message: "No, I'm not interested.",
					Title:   "No, thanks",
				},
			},
		},
		Message:  msg,
		Platform: "kommunicate",
	}

	fmt.Println(couponMessage)

	return couponMessage
}
