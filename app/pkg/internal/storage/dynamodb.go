package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
)

// DynamoDb is the DAO to manipulates the AWS DynamoDB
type DynamoDb struct {
	DbClient  *dynamodb.DynamoDB
	TableName string
	Timeout   time.Duration
}

// HasTables checks if there are tables in the database
func (db *DynamoDb) HasTables(ctx context.Context) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()

	log.Println("Executing HasTables")
	res, err := db.DbClient.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return false, err
	}

	return len(res.TableNames) > 0, nil
}

// CreateTables creates a new database structure
func (db *DynamoDb) CreateTables(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()
	log.Println("Executing CreateTables")
	params := &dynamodb.CreateTableInput{
		TableName: aws.String(db.TableName),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("UUID"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("UUID"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
	}

	_, err := db.DbClient.CreateTable(params)
	if err != nil {
		return err
	}
	return nil
}

// Create create a new votable in the database
func (db *DynamoDb) Create(ctx context.Context, data *entities.Voteable) (*entities.CreateVoteableResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()
	item, err := dynamodbattribute.MarshalMap(data)
	if err != nil {
		return nil, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item:      item,
	}

	_, err = db.DbClient.PutItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return &entities.CreateVoteableResponse{
		UUID: data.UUID,
	}, nil
}

// Update adds a new vote in a votable answer in the database
func (db *DynamoDb) Update(ctx context.Context, data *entities.CastVoteRequest) (*entities.CastVoteResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()
	answer, err := db.getAnswerByIndex(ctx, data.UUID)
	if err != nil {
		return nil, err
	}

	if data.AnswerIndex > int64(len(answer.Answers)) {
		return nil, errors.New("Answer index not found")
	}

	votes := 1
	if actVotes, ok := answer.Votes[fmt.Sprint(data.AnswerIndex)]; ok {
		votes = int(actVotes) + 1
	}

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(data.UUID)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#index": aws.String(fmt.Sprint(data.AnswerIndex)),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":answer_votes": {N: aws.String(fmt.Sprint(votes))},
		},
		UpdateExpression: aws.String("set Votes.#index = :answer_votes"),
	}

	_, err = db.DbClient.UpdateItemWithContext(ctx, input)
	if err != nil {
		return nil, err
	}

	return &entities.CastVoteResponse{
		Answer:      answer.Answers[data.AnswerIndex],
		AnswerVotes: int64(votes),
	}, nil
}

// List lists votables from the database using or not pagination
func (db *DynamoDb) List(ctx context.Context, data *entities.ListVoteableRequest) (*entities.ListVoteableResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()

	params := &dynamodb.ScanInput{
		TableName: aws.String(db.TableName),
	}

	if data.PageSize > 0 {
		params.Limit = aws.Int64(data.PageSize)
	}

	if data.Page != "" {
		esk, err := getPage(data.Page)
		if err != nil {
			return nil, err
		}
		params.ExclusiveStartKey = esk
	}

	res := []entities.Voteable{}
	resp, err := db.DbClient.Scan(params)
	if err != nil {
		return nil, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &res)
	if err != nil {
		return nil, err
	}

	nextPage := ""
	if resp.LastEvaluatedKey != nil {
		nextPage = getNextPage(resp.LastEvaluatedKey)
	}

	previousPage := getPreviousPage(res)

	response := &entities.ListVoteableResponse{
		Page:         data.Page,
		NextPage:     nextPage,
		PreviousPage: previousPage,
		Votables:     res,
	}

	return response, err
}

func (db *DynamoDb) getAnswerByIndex(ctx context.Context, uuid string) (entities.Voteable, error) {
	ctx, cancel := context.WithTimeout(ctx, db.Timeout)
	defer cancel()
	input := &dynamodb.GetItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]*dynamodb.AttributeValue{
			"UUID": {S: aws.String(uuid)},
		},
	}

	res, err := db.DbClient.GetItemWithContext(ctx, input)
	if err != nil {
		return entities.Voteable{}, err
	}

	if res.Item == nil {
		return entities.Voteable{}, errors.New("Not Found")
	}

	item := entities.Voteable{}
	err = dynamodbattribute.UnmarshalMap(res.Item, &item)
	if err != nil {
		return entities.Voteable{}, err
	}
	return item, nil
}
