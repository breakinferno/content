package content_storage

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"testing"
	"time"
)

// go test -run="TestInitDB"
func TestInitDB(t *testing.T) {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := initDB()
	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		t.Fatal(err)
	}
	t.Log(client)
}

