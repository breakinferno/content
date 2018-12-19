package content_storage

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

// DBInitialOptions
type DBInitialOptions struct {
	username string
	password string
}

func initDB(options ...*DBInitialOptions) (*mongo.Client, error) {
	var err error
	var client *mongo.Client

	client, err = mongo.NewClient("mongodb://localhost:27017")

	if err != nil {
		return nil, err
	}

	return client, nil
}

