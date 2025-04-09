package distribution

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/iypetrov/go-cr/logger"
)

type Storage struct {
	storageClient *s3.Client
	log           logger.Logger
}

func NewStorage(awsCfg aws.Config, log logger.Logger) *Storage {
	s3Client := s3.NewFromConfig(awsCfg)
	return &Storage{
		storageClient: s3Client,
		log:           log,
	}
}
