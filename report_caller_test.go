package logrushook_test

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook"
)

func ExampleReportCallerLogrusHook() {
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	//logrus.AddHook(
	//	logrushook.NewReportCallerLogrusHook(
	//		[]logrus.Level{logrus.ErrorLevel, logrus.WarnLevel},
	//		logrushook.DefaultPathHandler,
	//	),
	//)

	dir, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.AddHook(
		logrushook.NewReportCallerLogrusHook(
			[]logrus.Level{logrus.ErrorLevel, logrus.WarnLevel},
			func(path string) string {
				return strings.Replace(path, dir+"/", "", -1)
			},
		),
	)

	logrus.Error("error")
}
