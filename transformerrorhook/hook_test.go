package transformerrorhook

import (
	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

func Example() {
	var (
		notFoundCode            gocode.Code = "NotFound"
		internalServerErrorCode gocode.Code = "InternalServerError"
	)

	hook := New(logrus.WarnLevel)

	hook.ExcludedCodes([]gocode.Code{internalServerErrorCode})

	hook.DeleteErrorKey()

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(hook)

	log.WithError(notFoundCode).Error(notFoundCode)                       // WarnLevel
	log.WithError(internalServerErrorCode).Error(internalServerErrorCode) // ErrorLevel
}
