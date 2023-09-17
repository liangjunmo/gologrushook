package logrushook_test

import (
	"fmt"
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
	//		"file",
	//		logrushook.DefaultPathHandler,
	//	),
	//)

	dir, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}

	println(dir)

	logrus.AddHook(
		logrushook.NewReportCallerLogrusHook(
			[]logrus.Level{logrus.ErrorLevel, logrus.WarnLevel},
			"file",
			func(path string, line int) string {
				return fmt.Sprintf("%s:%d", strings.Replace(path, dir+"/", "", -1), line)
			},
		),
	)

	logrus.Error("error")
}
