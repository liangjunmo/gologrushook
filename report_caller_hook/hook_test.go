package report_caller_hook

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestHook(t *testing.T) {
	var (
		fieldKey = "location"
		location = "path:line"
		buffer   bytes.Buffer
		fields   logrus.Fields
	)

	hook := New([]logrus.Level{logrus.ErrorLevel}, fieldKey)

	hook.SetLocationHandler(func(fileAbsolutePath string, line int) string {
		return location
	})

	log := logrus.New()

	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(&buffer)

	log.AddHook(hook)

	log.Error("message")

	err := json.Unmarshal(buffer.Bytes(), &fields)
	require.Nil(t, err)
	require.Equal(t, location, fields[fieldKey])
}
