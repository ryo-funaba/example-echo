package dynamodb

import "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"

type DynamodbSession interface {
	GetClient() dynamodbiface.DynamoDBAPI
}
