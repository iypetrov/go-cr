package distribution

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/iypetrov/go-cr/logger"
)

type Metadata struct {
	metadataClient *dynamodb.Client
	log            logger.Logger
}

func NewMetadata(awsCfg aws.Config, log logger.Logger) *Metadata {
	dynamodbClient := dynamodb.NewFromConfig(awsCfg)
	return &Metadata{
		metadataClient: dynamodbClient,
		log:            log,
	}
}
