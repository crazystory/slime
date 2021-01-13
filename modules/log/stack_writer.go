package log

import (
	"github.com/spf13/viper"
)

var (
	_ Writer = &StackWriter{}
)

type StackWriter struct {
	channels []Writer
}

func (r *StackWriter) Write(p []byte) (n int, err error) {
	for _, w := range r.channels {
		n, err = w.Write(p)
	}
	return
}

type StackWriterConfig struct {
	Channels []string `mapstructure:"channels"`
}

func newStackDriver(c StackWriterConfig) (*StackWriter, error) {
	channels := make([]Writer, 0)
	for _, name := range c.Channels {
		w, err := createChannel(name)
		if err != nil {
			return nil, err
		}
		channels = append(channels, w)
	}

	return &StackWriter{
		channels: channels,
	}, nil
}

func init() {
	Register(`stack`, func(v *viper.Viper) (Writer, error) {
		c := StackWriterConfig{
			Channels: make([]string, 0),
		}
		if err := v.Unmarshal(&c); err != nil {
			return nil, err
		}

		return newStackDriver(c)
	})
}
