package logrushook

import (
	"errors"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

type transErrorToWarnLogrusHook struct {
	errCodes []gocode.Code
}

func NewTransErrorToWarnLogrusHook(errCodes []gocode.Code) logrus.Hook {
	return &transErrorToWarnLogrusHook{
		errCodes: errCodes,
	}
}

func (hook *transErrorToWarnLogrusHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *transErrorToWarnLogrusHook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)

	if !ok || err == nil {
		return nil
	}

	for _, code := range hook.errCodes {
		if errors.Is(err, code) {
			return nil
		}
	}

	entry.Level = logrus.WarnLevel

	return nil
}
