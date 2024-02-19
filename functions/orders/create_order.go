package orders

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
)

var busName = os.Getenv("EVENT_BUS_NAME")

type CreateOrderRequest struct {
	Amount    int64    `json:"amount"`
	LineItems []string `json:"line_items"`
}

func CreateOrderHandler(request events.APIGatewayProxyRequest) error {
	// 1. Unmarshal request
	var req CreateOrderRequest
	if err := json.Unmarshal([]byte(request.Body), &req); err != nil {
		return err
	}

	// 2. Set up EventBridge client
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}
	client := cloudwatchevents.NewFromConfig(cfg)

	// 3. Marshal order into JSON again
	b, err := json.Marshal(req)
	if err != nil {
		return err
	}

	// 4. Publish events to EventBridge
	_, err = client.PutEvents(context.Background(),
		&cloudwatchevents.PutEventsInput{
			Entries: []types.PutEventsRequestEntry{
				{
					EventBusName: aws.String(busName),
					Source:       aws.String("create_order_fn"),

					DetailType: aws.String("order.created"),
					Detail:     aws.String(string(b)),
				},
			},
		})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	lambda.Start(CreateOrderHandler)
}
