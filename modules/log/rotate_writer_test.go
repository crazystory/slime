package log

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestRotate(t *testing.T) {
	c := `log:
  default: stack

  channels:
    stack:
      writer: stack
      channels: ["rotate"]

    rotate:
      writer: rotate
      pattern: "slime-2006-01-02.log"`

	viper.SetConfigType(`yaml`)
	assert.NoError(t, viper.ReadConfig(bytes.NewBufferString(c)))

	channel, err := createChannel(`rotate`)
	assert.NoError(t, err)
	_, err = channel.Write([]byte(`rotate`))

	assert.NoError(t, err)

	filename := time.Now().Format(`slime-2006-01-02.log`)
	defer os.RemoveAll(filename)

	content, err := ioutil.ReadFile(filename)
	assert.NoError(t, err)
	assert.Equal(t, `rotate`, string(content))
}
