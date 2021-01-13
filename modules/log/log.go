package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
)

var (
	Logger  *logrus.Logger
	creators = make(map[string]CreateWriter)
)

type CreateWriter func(v *viper.Viper) (Writer, error)

// register writer driver
func Register(writer string, fn CreateWriter) {
	creators[writer] = fn
}

func Init(channel string) error {
	if channel, err := createChannel(channel); err != nil {
		return err
	} else {
		Logger = logrus.New()
		Logger.Out = channel
		return nil
	}
}

type Writer interface {
	io.Writer
}

func createChannel(name string) (Writer, error) {
	v := viper.Sub(fmt.Sprintf(`log.channels.%s`, name))
	if v == nil {
		return nil, fmt.Errorf(`log [%s] is not defined`, name)
	}

	creator, ok := creators[v.GetString(`writer`)]
	if !ok {
		return nil, fmt.Errorf(`driver [%s] is not supported`, name)
	}
	return creator(v)
}
