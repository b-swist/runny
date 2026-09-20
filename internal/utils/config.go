package utils

type ConfigOption[C any] interface {
	Apply(*C)
}

type ConfigFunc[C any] func(*C)

func (f ConfigFunc[C]) Apply(cfg *C) {
	f(cfg)
}
