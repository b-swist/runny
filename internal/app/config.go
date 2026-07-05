package app

type modelConfig struct {
	prompt          string
	showDescription bool
}

type ModelOption interface {
	apply(*modelConfig)
}

type modelOptionFunc func(*modelConfig)

func (f modelOptionFunc) apply(cfg *modelConfig) {
	f(cfg)
}

func NewModelConfig(opts ...ModelOption) *modelConfig {
	cfg := &modelConfig{
		prompt:          "runny",
		showDescription: true,
	}

	for _, opt := range opts {
		opt.apply(cfg)
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
