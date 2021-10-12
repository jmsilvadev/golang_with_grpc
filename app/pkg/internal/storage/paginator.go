package storage

import (
	"encoding/base64"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
)

func getNextPage(esk map[string]*dynamodb.AttributeValue) string {
	lastkey := map[string]string{}
	dynamodbattribute.UnmarshalMap(esk, &lastkey)
	lk, err := json.Marshal(lastkey)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(lk)
}

func getPreviousPage(res []entities.Voteable) string {
	if len(res) > 0 {
		lastkey := &entities.CreateVoteableResponse{
			UUID: res[0].UUID,
		}
		lk, err := json.Marshal(lastkey)
		if err != nil {
			return ""
		}
		return base64.StdEncoding.EncodeToString(lk)
	}
	return ""
}

func getPage(page string) (map[string]*dynamodb.AttributeValue, error) {
	b, err := base64.StdEncoding.DecodeString(page)
	if err != nil {
		return nil, err
	}
	esk := map[string]string{}
	json.Unmarshal(b, &esk)
	return dynamodbattribute.MarshalMap(esk)
}
