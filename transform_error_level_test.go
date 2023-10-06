package logrushook_test

import (
	"testing"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"

	"github.com/liangjunmo/logrushook"
)

func TestTransformErrorLevelLogrusHookWithoutExcludeCodes(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	log.AddHook(
		logrushook.NewTransformErrorLevelLogrusHook(
			logrus.WarnLevel,
			nil,
			true,
		),
	)

	var errorCode gocode.Code = "Error"
	log.WithError(errorCode).Error("transform to warn level")
}

func TestTransformErrorLevelLogrusHookWithExcludeCodes(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var errorCode gocode.Code = "Error"

	log.AddHook(
		logrushook.NewTransformErrorLevelLogrusHook(
			logrus.WarnLevel,
			[]gocode.Code{errorCode},
			false,
		),
	)

	log.WithError(errorCode).Error("not transform to warn level")
}
