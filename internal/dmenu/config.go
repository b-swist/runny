package dmenu

import "github.com/b-swist/runny/internal/utils"

type dmenuConfig struct {
	delimeter string
	input     string
}

type DmenuOption = utils.ConfigOption[dmenuConfig]
type dmenuOptionFunc = utils.ConfigFunc[dmenuConfig]

func WithDelimeter(d string) DmenuOption {
	return dmenuOptionFunc(func(cfg *dmenuConfig) {
		cfg.delimeter = d
	})
}

func WithInput(input string) DmenuOption {
	return dmenuOptionFunc(func(cfg *dmenuConfig) {
		cfg.input = input
	})
}
