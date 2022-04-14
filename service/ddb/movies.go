package ddb

import (
	"github.com/FerVT/movies/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const (
	idKey = "id"
)

func (db *movies) GetAllMovies() ([]*model.Movie, error) {
	scanInput := &dynamodb.ScanInput{
		TableName: &db.tableName,
	}

	scanOutput, err := db.client.Scan(scanInput)
	if err != nil {
		return nil, err
	}

	var movies []*model.Movie
	err = dynamodbattribute.UnmarshalListOfMaps(scanOutput.Items, &movies)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (db *movies) GetMovieByID(movieId string) (*model.Movie, error) {
	getItemInput := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			idKey: {
				S: aws.String(movieId),
			},
		},
		TableName: &db.tableName,
	}

	getItemOutput, err := db.client.GetItem(getItemInput)
	if err != nil {
		return nil, err
	}

	if len(getItemOutput.Item) == 0 {
		return nil, nil
	}

	var movie *model.Movie
	err = dynamodbattribute.UnmarshalMap(getItemOutput.Item, &movie)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

func (db *movies) CreateMovies(movies []*model.Movie) ([]*model.Movie, error) {
	writeRequests := make([]*dynamodb.WriteRequest, len(movies))

	for i, m := range movies {
		movieMap, err := dynamodbattribute.MarshalMap(m)
		if err != nil {
			return nil, err
		}

		writeRequests[i] = &dynamodb.WriteRequest{
			PutRequest: &dynamodb.PutRequest{
				Item: movieMap,
			},
		}
	}

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			db.tableName: writeRequests,
		},
	}

	_, err := db.client.BatchWriteItem(batchWriteItemInput)
	if err != nil {
		return nil, err
	}

	return movies, nil
}

func (db *movies) DeleteMoviesByIds(moviesIds []string) error {
	writeRequests := make([]*dynamodb.WriteRequest, len(moviesIds))

	for i, mId := range moviesIds {
		writeRequests[i] = &dynamodb.WriteRequest{
			DeleteRequest: &dynamodb.DeleteRequest{
				Key: map[string]*dynamodb.AttributeValue{
					idKey: {
						S: aws.String(mId),
					},
				},
			},
		}
	}

	batchWriteItemInput := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{
			db.tableName: writeRequests,
		},
	}

	_, err := db.client.BatchWriteItem(batchWriteItemInput)
	if err != nil {
		return err
	}

	return nil
}
