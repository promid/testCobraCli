package cmd

import (
	"github.com/spf13/pflag"
)

type Options struct {
	Test string
}

func (option *Options) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&option.Test, "test", "0", "")
}

func NewRunOptions() *Options {
	return &Options{}
}
