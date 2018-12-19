package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) UpdateContent(ctx context.Context, r *content.UpdateContentRequest) (*content.UpdateContentResponse, error) {
	return nil, errors.New("not yet implemented")
}
