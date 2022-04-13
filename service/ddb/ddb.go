package ddb

import (
	"github.com/FerVT/movies/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Movies struct {
	tableName string
	client    ddbClient
}

type ddbClient interface {
	DescribeTable(input *dynamodb.DescribeTableInput) (*dynamodb.DescribeTableOutput, error)
}

func NewMovies(cfg *config.Config) (*Movies, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.AWSRegion),
		Endpoint:    aws.String(cfg.AWSHost),
		Credentials: credentials.NewStaticCredentials(cfg.AWSID, cfg.AWSSecret, cfg.AWSToken),
	})
	if err != nil {
		return nil, err
	}

	moviesDB := &Movies{
		tableName: cfg.MoviesTableName,
		client:    dynamodb.New(sess),
	}

	err = moviesDB.testConnection()
	if err != nil {
		return nil, err
	}

	return moviesDB, nil
}

func (db *Movies) testConnection() error {
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String(db.tableName),
	}

	_, err := db.client.DescribeTable(req)
	return err
}
