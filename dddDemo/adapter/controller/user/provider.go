package user

import (
    "AF-Excel/domain/service"
    "AF-Excel/interface/user"
    "af-go-frame/core/container"
)

type UserProvider struct {
    container.ServiceProvider
    c container.Container
}


func (sp *UserProvider) Name() string {
    return user.UserKey
}

func (sp *UserProvider) Register(c container.Container) container.NewInstance {
    return service.NewUserService
}

func (sp *UserProvider) IsDefer() bool {
    return true
}

func (sp *UserProvider) Params(c container.Container) []interface{} {
    return []interface{}{c}
}

func (sp *UserProvider) Boot(c container.Container) error {
    return nil
}
