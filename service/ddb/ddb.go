package ddb

import (
	"github.com/FerVT/movies/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type movies struct {
	tableName string
	client    ddbClient
}

type ddbClient interface {
	DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
	Scan(input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
	GetItem(input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error)
	BatchWriteItem(input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error)
}

func NewMovies(cfg *config.Config) (*movies, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      &cfg.AWSRegion,
		Endpoint:    &cfg.AWSHost,
		Credentials: credentials.NewStaticCredentials(cfg.AWSID, cfg.AWSSecret, cfg.AWSToken),
	})
	if err != nil {
		return nil, err
	}

	moviesDB := &movies{
		tableName: cfg.MoviesTableName,
		client:    dynamodb.New(sess),
	}

	err = moviesDB.testConnection()
	if err != nil {
		return nil, err
	}

	return moviesDB, nil
}

func (db *movies) testConnection() error {
	req := &dynamodb.DescribeTableInput{
		TableName: &db.tableName,
	}

	_, err := db.client.DescribeTable(req)
	return err
}
