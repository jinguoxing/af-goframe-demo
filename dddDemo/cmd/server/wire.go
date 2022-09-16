// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
    af_go_frame "af-go-frame"
    "af-goframe-demo/dddDemo/adapter/controller"
    "af-goframe-demo/dddDemo/domain"
    "af-goframe-demo/dddDemo/infrastructure/conf"
    "af-goframe-demo/dddDemo/infrastructure/repository"
    "github.com/google/wire"
)


func InitApp(*conf.Server, *conf.Data) (*af_go_frame.App, func(), error) {
   // panic(wire.Build(controller.ProviderSet, repository.RepositoryProviderSet, domain.DomainProviderSet, newApp))
    panic(wire.Build(controller.HttpProviderSet,controller.ServiceProviderSet,domain.DomainProviderSet,repository.RepositoryProviderSet,newApp))
}
