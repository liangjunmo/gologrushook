package transform_error_level_hook

import (
	"fmt"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

func Example() {
	var (
		internalServerErrorCode gocode.Code = "InternalServerError"
		notFoundCode            gocode.Code = "NotFound"
	)

	hook := New(logrus.WarnLevel)

	hook.DeleteErrorKey()

	hook.SetExcludedCodes([]gocode.Code{internalServerErrorCode})

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(hook)

	err := fmt.Errorf("err")
	log.WithError(err).Error(err) // ErrorLevel

	log.WithError(internalServerErrorCode).Error(internalServerErrorCode) // ErrorLevel

	log.WithError(notFoundCode).Error(notFoundCode) // WarnLevel
}
