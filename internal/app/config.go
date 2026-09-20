package app

import "github.com/b-swist/runny/internal/utils"

type modelConfig struct {
	prompt          string
	showDescription bool
}

type ModelOption = utils.ConfigOption[modelConfig]
type modelOptionFunc = utils.ConfigFunc[modelConfig]

func NewModelConfig(opts ...ModelOption) *modelConfig {
	cfg := &modelConfig{
		prompt:          "runny",
		showDescription: true,
	}

	for _, opt := range opts {
		opt.Apply(cfg)
	}

	return cfg
}

func WithPrompt(prompt string) ModelOption {
	return modelOptionFunc(func(cfg *modelConfig) {
		cfg.prompt = prompt
	})
}

func WithDesctiption(show bool) ModelOption {
	return modelOptionFunc(func(cfg *modelConfig) {
		cfg.showDescription = show
	})
}
