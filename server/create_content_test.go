package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestCreateContent(t *testing.T) {
	ctx := context.Background()
	req := &content.CreateContentRequest{}

	res, err := cli.CreateContent(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
