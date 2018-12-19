package server

import (
	"testing"

	"context"
	"github.com/stretchr/testify/assert"
	"github.com/sundogrd/content"
)

func TestDeleteContent(t *testing.T) {
	ctx := context.Background()
	req := &content.DeleteContentRequest{}

	res, err := cli.DeleteContent(ctx, req)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
