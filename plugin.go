package app

import (
	"github.com/roadrunner-server/http/v4/common"
	"go.uber.org/zap"
)

const pluginName = "app"

type Logger interface {
	NamedLogger(name string) *zap.Logger
}

type Plugin struct {
	log    *zap.Logger
	config *Config
}

func (p *Plugin) Init(cfg common.Configurer, log Logger) error {
	p.log = log.NamedLogger(pluginName)

	p.config = &Config{}
	err := cfg.UnmarshalKey(pluginName, &p.config)
	if err != nil {
		return err
	}

	return nil
}

func (p *Plugin) Name() string {
	return pluginName
}

func (p *Plugin) RPC() any {
	return &RPC{
		log: p.log,
		cfg: p.config,
	}
}
