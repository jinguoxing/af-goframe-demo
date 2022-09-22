package repository

import (
	"af-goframe-demo/dddDemo/infrastructure/conf"
	"github.com/google/wire"
)

// RepositoryProviderSet is data providers.
var RepositoryProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data) (*Data, func(), error) {
	cleanup := func() {
		// log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}
