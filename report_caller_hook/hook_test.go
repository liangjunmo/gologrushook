package report_caller_hook

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func Example() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	hook := New([]logrus.Level{logrus.ErrorLevel})

	hook.SetKey("location")

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

	log.Error("message")
}
