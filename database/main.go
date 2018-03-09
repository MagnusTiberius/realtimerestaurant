package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const CREATING = "CREATING..."

func main() {

	config := &aws.Config{
		Region: aws.String("us-central1-a"),
	}

	svc := dynamodb.New(config)

	_, err := tables.CreateReservationTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)

}
