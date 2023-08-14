package logrushook_test

import (
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook"
)

func ExampleReportCallerLogrusHook() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//logrus.AddHook(logrushook.NewReportCallerLogrusHook([]logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}, logrushook.DefaultPathHandler))

	logrus.AddHook(logrushook.NewReportCallerLogrusHook([]logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}, func(path string) string {
		return strings.Replace(path, dir+"/", "", -1)
	}))

	logrus.Error("error")
}
