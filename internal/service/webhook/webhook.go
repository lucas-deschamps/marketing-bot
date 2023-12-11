package webhook

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"

	. "github.com/lucas-deschamps/marketing-bot/internal/model"
)

type webhookService struct{}

func NewService() *webhookService {
	return &webhookService{}
}

func (s *webhookService) ProcessMessage(c *fiber.Ctx) error {
	var ProjectID string = os.Getenv("PROJECT_ID")

	fmt.Printf("\nProjectID: %s", ProjectID)

	fmt.Println(string(c.Body()))

	jsonReq := new(DialogFlowRequest)

	if err := c.BodyParser(jsonReq); err != nil {
		fmt.Println("Error when parsing incoming body request: ", err)
		return c.SendStatus(500)
	}

	fmt.Println(jsonReq.QueryResult.Action)

	// var bot Chatbot = chatbot.New()

	test := RichResponse{
		FulfillmentMessages: []FulfillmentMessages{
			FulfillmentMessages{
				Card: Card{
					Title:    "Test button",
					Subtitle: "Test subtitle",
					ImageURI: "https://www.stickersdevs.com.br/wp-content/uploads/2022/01/gopher-adesivo-sticker.png",
					Buttons: []Button{
						Button{
							Text:     "Buttonnnn",
							Postback: "https://www.google.com",
						},
						Button{
							Text:     "Buttonnnn2",
							Postback: "https://www.google.com",
						},
					},
				},
			},
		},
	}

	// fmt.Println(test)

	// response := bot.Message("Webhooked")

	// b, _ := json.Marshal(test)

	// fmt.Println(string(b))

	return c.JSON(test)
}
