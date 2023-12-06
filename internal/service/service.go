package service

import "context"

type Service struct {
	pollers int64
	context context.Context
	cancel  context.CancelFunc
}

func NewService(ctx context.Context) (*Service, error) {
	ctx, cancel := context.WithCancel(ctx)
	s := &Service{
		context: ctx,
		cancel:  cancel,
		pollers: 0,
	}
	return s, nil
}
