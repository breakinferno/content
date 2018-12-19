package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) ListContents(ctx context.Context, r *content.ListContentsRequest) (*content.ListContentsResponse, error) {
	return nil, errors.New("not yet implemented")
}
