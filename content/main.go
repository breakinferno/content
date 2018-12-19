package main

import (
	_ "net/http/pprof"

	"github.com/lileio/fromenv"
	"github.com/lileio/lile"
	"github.com/lileio/logr"
	"github.com/lileio/pubsub"
	"github.com/lileio/pubsub/middleware/defaults"
	"github.com/sundogrd/content"
	"github.com/sundogrd/content/content/cmd"
	"github.com/sundogrd/content/server"
	"google.golang.org/grpc"
)

func main() {
	logr.SetLevelFromEnv()
	s := &server.ContentServer{}

	lile.Name("content")
	lile.Server(func(g *grpc.Server) {
		content.RegisterContentServer(g, s)
	})

	pubsub.SetClient(&pubsub.Client{
		ServiceName: lile.GlobalService().Name,
		Provider:    fromenv.PubSubProvider(),
		Middleware:  defaults.Middleware,
	})

	cmd.Execute()
}
