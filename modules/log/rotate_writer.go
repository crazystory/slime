package log

import (
	"github.com/spf13/viper"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var (
	_ io.Writer = &RotateWriter{}
)

type RotateWriter struct {
	mutex   sync.RWMutex
	pattern string
}

func (r *RotateWriter) Write(p []byte) (n int, err error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	filename := r.filename()

	// create file
	dirname := filepath.Dir(filename)

	if err := os.MkdirAll(dirname, 0755); err != nil {
		return 0, err
	}
	fp, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}

	defer fp.Close()
	return fp.Write(p)
}

func (r *RotateWriter) filename() string {
	return time.Now().Format(r.pattern)
}

type RotateConfig struct {
	Pattern string `mapstructure:"pattern"`
}

func newRotateWriter(config RotateConfig) *RotateWriter {
	return &RotateWriter{
		pattern: config.Pattern,
		mutex:   sync.RWMutex{},
	}
}

func init() {
	Register(`rotate`, func(v *viper.Viper) (Writer, error) {
		c := RotateConfig{}
		if err := v.Unmarshal(&c); err != nil {
			return nil, err
		}
		return newRotateWriter(c), nil
	})
}
