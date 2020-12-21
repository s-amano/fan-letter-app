package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// SurveyFanLetter - ファンレター内容を表現する構造体です
type SurveyFanLetter struct {
	ID   string `dynamodbav:"id" json:"id"`
	Text string `dynamodbav:"text" json:"text"`
}

// Scan はDynamoDBから全件取得する関数です
func Scan() []SurveyFanLetter {
	var fanLetters []SurveyFanLetter = []SurveyFanLetter{}
	session, err := session.NewSession()
	conn := dynamodb.New(session)
	s := &dynamodb.ScanInput{
		TableName: aws.String("fan-letter-app-table"),
	}
	resp, err := conn.Scan(s)

	if err != nil {
		fmt.Println(err.Error())
		return fanLetters
	}

	// 取得したファンレターそれぞれに対して処理
	for _, scanedFanLetter := range resp.Items {
		var fanLetter SurveyFanLetter
		_ = dynamodbattribute.UnmarshalMap(scanedFanLetter, &fanLetter)
		fanLetters = append(fanLetters, fanLetter)
	}
	fmt.Printf("fanletterの中身だよ %v", fanLetters)
	return fanLetters

}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	allFanLetters := Scan()

	j, err := json.Marshal(allFanLetters)

	if err != nil {
		fmt.Println(err.Error())
	}

	return events.APIGatewayProxyResponse{
		Body:       string(j),
		StatusCode: 200,
	}, nil

}

func main() {
	Scan()
	lambda.Start(handler)
}
