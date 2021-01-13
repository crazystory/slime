package log

import (
	"github.com/spf13/viper"
	"os"
)

var (
	_ Writer = &StdWriter{}
)

type StdWriter struct {
}

func (s StdWriter) Write(p []byte) (n int, err error) {
	return os.Stderr.Write(p)
}

func init() {
	Register(`std`, func(_ *viper.Viper) (Writer, error) {
		return &StdWriter{}, nil
	})
}
