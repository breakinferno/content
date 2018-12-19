package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	req := &content.GetRequest{}

	res, err := cli.Get(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
