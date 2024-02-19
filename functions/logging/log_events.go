package logging

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/davecgh/go-spew/spew"
)

func LogHandler(request events.CloudWatchEvent) error {
	fmt.Printf("received event of type %q\n", request.DetailType)
	spew.Dump(request)
	return nil
}

func main() {
	lambda.Start(LogHandler)
}
