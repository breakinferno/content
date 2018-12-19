package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) GetContent(ctx context.Context, r *content.GetContentRequest) (*content.GetContentResponse, error) {
	return nil, errors.New("not yet implemented")
}
