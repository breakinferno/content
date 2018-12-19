package server

import (
	"errors"

	"context"
	"github.com/sundogrd/content"
)

func (s ContentServer) DeleteContent(ctx context.Context, r *content.DeleteContentRequest) (*content.DeleteContentResponse, error) {
	return nil, errors.New("not yet implemented")
}
