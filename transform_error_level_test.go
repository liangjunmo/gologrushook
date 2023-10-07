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
			true, // delete err key
		),
	)

	var notFoundCode gocode.Code = "NotFound"

	log.WithError(notFoundCode).Error("TestTransformErrorLevelLogrusHookWithoutExcludeCodes") // transform to warn level
}

func TestTransformErrorLevelLogrusHookWithExcludeCodes(t *testing.T) {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote:    true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var notFoundCode gocode.Code = "NotFound"

	log.AddHook(
		logrushook.NewTransformErrorLevelLogrusHook(
			logrus.WarnLevel,
			[]gocode.Code{notFoundCode},
			false, // not delete err key
		),
	)

	log.WithError(notFoundCode).Error("TestTransformErrorLevelLogrusHookWithExcludeCodes") // not transform to warn level
}
