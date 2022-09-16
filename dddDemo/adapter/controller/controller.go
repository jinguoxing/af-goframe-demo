package controller

import (
    "af-goframe-demo/dddDemo/adapter/controller/greeter"
    "github.com/google/wire"
)

// ProviderSet is server providers.
var HttpProviderSet = wire.NewSet(NewHttpServer)

var ServiceProviderSet = wire.NewSet(greeter.NewGreeterService)
