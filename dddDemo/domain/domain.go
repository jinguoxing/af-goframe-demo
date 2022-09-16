package domain

import (
    "af-goframe-demo/dddDemo/domain/greeter"
    "github.com/google/wire"
)

// ProviderSet is biz providers.
var DomainProviderSet = wire.NewSet(greeter.NewGreeterUsecase)
