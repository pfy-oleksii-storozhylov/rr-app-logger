package app

import (
	"go.uber.org/zap"
)

const (
	name = "app"
)

type Plugin struct {
	log *zap.Logger
}

func (p *Plugin) Init(log *zap.Logger) error {
	p.log = new(zap.Logger)
	*p.log = *log

	return nil
}

func (p *Plugin) Name() string {
	return name
}

func (p *Plugin) RPC() any {
	return &RPC{
		log: p.log,
	}
}
