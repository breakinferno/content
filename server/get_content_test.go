package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestGetContent(t *testing.T) {
	ctx := context.Background()
	req := &content.GetContentRequest{}

	res, err := cli.GetContent(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
