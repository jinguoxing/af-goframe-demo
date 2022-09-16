package repository

import (
    "af-goframe-demo/dddDemo/domain/greeter"
    "af-goframe-demo/dddDemo/infrastructure/conf"
    "context"
)

type greeterRepo struct {
    data *conf.Data
    //log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *conf.Data) greeter.GreeterRepo {
    return &greeterRepo{
        data: data,
       // log:  log.NewHelper(logger),
    }
}

func (r *greeterRepo) Save(ctx context.Context, g *greeter.Greeter) (*greeter.Greeter, error) {
    return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *greeter.Greeter) (*greeter.Greeter, error) {
    return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*greeter.Greeter, error) {
    return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*greeter.Greeter, error) {
    return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*greeter.Greeter, error) {
    return nil, nil
}
