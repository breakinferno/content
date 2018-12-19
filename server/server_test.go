package server

import (
	"os"
	"testing"

	"google.golang.org/grpc"

	"github.com/lileio/lile"
	"github.com/sundogrd/content"
)

var s = ContentServer{}
var cli content.ContentClient

func TestMain(m *testing.M) {
	impl := func(g *grpc.Server) {
		content.RegisterContentServer(g, s)
	}

	gs := grpc.NewServer()
	impl(gs)

	addr, serve := lile.NewTestServer(gs)
	go serve()

	cli = content.NewContentClient(lile.TestConn(addr))

	os.Exit(m.Run())
}
