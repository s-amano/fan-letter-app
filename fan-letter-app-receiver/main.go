package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const (
	tableName = "fan-letter-app-table"
)

// FanLetterRequest は送信されたバイト配列をつめるためのファンレター内容を表す構造体です
type FanLetterRequest struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	From string `json:"from"`
}

// PutDynamoDB はDynamoDBへitemをputする関数です
func PutDynamoDB(request events.APIGatewayProxyRequest) []byte {
	reqBody := request.Body
	jsonBytes := ([]byte)(reqBody)
	fanletterReq := new(FanLetterRequest)

	// FanLetterRequest構造体に格納
	if err := json.Unmarshal(jsonBytes, fanletterReq); err != nil {
		fmt.Println("[ERROR]", err)
	}

	id := fanletterReq.ID
	text := fanletterReq.Text
	from := fanletterReq.From

	fmt.Printf("id %v, text %v from %v\n", id, text, from)

	fanletter := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(id),
		},
		"text": {
			S: aws.String(text),
		},
		"from": {
			S: aws.String(from),
		},
	}

	session, err := session.NewSession()
	conn := dynamodb.New(session)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	p := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      fanletter,
	}

	fmt.Printf("pの値 %v\n", p)

	resp, err := conn.PutItem(p)
	fmt.Printf("respだよ %v", resp)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	return jsonBytes
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	jsonBytes := PutDynamoDB(request)
	return events.APIGatewayProxyResponse{
		Body: string(jsonBytes),
		Headers: map[string]string{
			"Content-Type":                 "text/html",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "POST",
		},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
