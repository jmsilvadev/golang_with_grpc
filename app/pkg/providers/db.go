package providers

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/entities"
	"github.com/jmsilvadev/golangtechtask/pkg/internal/storage"
)

// Storage is the interface that the database provider must be implement to use the repository pattern
type Storage interface {
	HasTables(ctx context.Context) (bool, error)
	CreateTables(ctx context.Context) error
	Create(ctx context.Context, data *entities.Voteable) (*entities.CreateVoteableResponse, error)
	Update(ctx context.Context, data *entities.CastVoteRequest) (*entities.CastVoteResponse, error)
	List(ctx context.Context, data *entities.ListVoteableRequest) (*entities.ListVoteableResponse, error)
}

//NewDynamoDB gets a DynamoDB Instance
func NewDynamoDB(ctx context.Context, conn *session.Session, duration time.Duration) *storage.DynamoDb {
	client := dynamodb.New(conn)
	db := &storage.DynamoDb{
		DbClient:  client,
		Timeout:   duration,
		TableName: "Voteables",
	}

	ok, err := db.HasTables(ctx)
	if err != nil {
		log.Println(err)
	}

	if !ok {
		err := db.CreateTables(ctx)
		if err != nil {
			log.Println(err)
		}
	}
	return db
}

//NewMockDB gets a MockDB Instance
func NewMockDB(ctx context.Context) *storage.MockDb {
	return &storage.MockDb{}
}
