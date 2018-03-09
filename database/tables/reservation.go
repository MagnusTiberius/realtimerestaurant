package reservation

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//CreateReservationTable table
func CreateReservationTable(svc *dynamodb.DynamoDB) (*dynamodb.CreateTableOutput, error) {
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Reservation"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	return svc.CreateTable(params)
}
