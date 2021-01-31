package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"net/url"
)

const (
	tableName = "fan-letter-app-table"
)

// DeleteDynamoDB はDynamoDBへitemをputする関数です
func DeleteDynamoDB(request events.APIGatewayProxyRequest) error {

	reqPath := request.PathParameters
	fmt.Printf("reqPath %v \n", reqPath)

	deleteID := reqPath["id"]
	fmt.Printf("deleteID %v \n", deleteID)

	deleteTextPara := reqPath["text"]
	fmt.Printf("deleteTextPara %v \n", deleteTextPara)

	deleteText, err := url.QueryUnescape(deleteTextPara)
	if err != nil {
		return err
	}

	fmt.Printf("deleteText %v \n", deleteText)

	fanletter := map[string]*dynamodb.AttributeValue{
		"id": {
			S: aws.String(deleteID),
		},
		"text": {
			S: aws.String(deleteText),
		},
	}

	session, err := session.NewSession()
	conn := dynamodb.New(session)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	p := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       fanletter,
	}

	resp, err := conn.DeleteItem(p)
	fmt.Printf("respだよ %v", resp)
	if err != nil {
		fmt.Println("[ERROR]", err)
	}

	return nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := DeleteDynamoDB(request)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "error!",
			StatusCode: 400, // Bad Request
		}, err
	}
	return events.APIGatewayProxyResponse{
		Body: "success!",
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "DELETE",
		},
		StatusCode: 200}, nil
}

func main() {
	lambda.Start(handler)
}
