package logrushook

import (
	"errors"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

type TransErrorToWarningLogrusHook struct {
	excludeCodes     []gocode.Code
	deleteErrorField bool
}

func NewTransErrorToWarningLogrusHook(excludeCodes []gocode.Code, deleteErrorField bool) logrus.Hook {
	return &TransErrorToWarningLogrusHook{
		excludeCodes:     excludeCodes,
		deleteErrorField: deleteErrorField,
	}
}

func (hook *TransErrorToWarningLogrusHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *TransErrorToWarningLogrusHook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)

	if !ok || err == nil {
		return nil
	}

	if hook.deleteErrorField {
		delete(entry.Data, logrus.ErrorKey)
	}

	for _, code := range hook.excludeCodes {
		if errors.Is(err, code) {
			return nil
		}
	}

	entry.Level = logrus.WarnLevel

	return nil
}
