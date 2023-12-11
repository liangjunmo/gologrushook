package reportcallerhook_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook/reportcallerhook"
)

func TestHook(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	hook := reportcallerhook.New([]logrus.Level{logrus.ErrorLevel})

	hook.SetKey("file")

	hook.SetLocationHandler(func(path string, line int) string {
		return fmt.Sprintf("%s:%d", strings.Replace(path, dir+"/", "", -1), line)
	})

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(hook)

	log.Error("test")
}
