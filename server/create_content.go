package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) CreateContent(ctx context.Context, r *content.CreateContentRequest) (*content.CreateContentResponse, error) {
	return nil, errors.New("not yet implemented")
}
