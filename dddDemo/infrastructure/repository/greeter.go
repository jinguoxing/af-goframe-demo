package repository

import (
	"af-goframe-demo/dddDemo/domain/greeter"
	"af-goframe-demo/dddDemo/infrastructure/conf"
	"context"
	"go.uber.org/zap"
)

type greeterRepo struct {
	data *conf.Data
	log  *zap.Logger
}

// NewGreeterRepo .
func NewGreeterRepo(data *conf.Data, log *zap.Logger) greeter.GreeterRepo {

	return &greeterRepo{
		data: data,
		log:  log,
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *greeter.Greeter) (*greeter.Greeter, error) {
	r.log.Info("Save")
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
