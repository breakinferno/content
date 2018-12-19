package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestListContents(t *testing.T) {
	ctx := context.Background()
	req := &content.ListContentsRequest{}

	res, err := cli.ListContents(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
