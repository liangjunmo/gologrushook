package logrushook_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook"
)

func TestReportCallerLogrusHookWithDefaultPathHandler(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(
		logrushook.NewReportCallerLogrusHook(
			[]logrus.Level{logrus.ErrorLevel},
			"file",
			logrushook.DefaultPathHandler,
		),
	)

	log.Error("error message")
}

func TestReportCallerLogrusHookWithCustomPathHandler(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	log.AddHook(
		logrushook.NewReportCallerLogrusHook(
			[]logrus.Level{logrus.ErrorLevel},
			"file",
			func(path string, line int) string {
				return fmt.Sprintf("%s:%d", strings.Replace(path, dir+"/", "", -1), line)
			},
		),
	)

	log.Error("error message")
}
