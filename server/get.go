package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) Get(ctx context.Context, r *content.GetRequest) (*content.GetResponse, error) {
	return nil, errors.New("not yet implemented")
}
