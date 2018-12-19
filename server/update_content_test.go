package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestUpdateContent(t *testing.T) {
	ctx := context.Background()
	req := &content.UpdateContentRequest{}

	res, err := cli.UpdateContent(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
