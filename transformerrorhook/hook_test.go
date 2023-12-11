package transformerrorhook_test

import (
	"testing"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook/transformerrorhook"
)

func TestHook(t *testing.T) {
	var (
		internalServerErrorCode gocode.Code = "InternalServerError"
		notFoundCode            gocode.Code = "NotFound"
	)

	hook := transformerrorhook.New(logrus.WarnLevel)

	hook.ExcludeCodes([]gocode.Code{internalServerErrorCode})
	hook.DeleteErrorKey()

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(hook)

	log.WithError(internalServerErrorCode).Error(internalServerErrorCode)
	log.WithError(notFoundCode).Error(notFoundCode)
}
