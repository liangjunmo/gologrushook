package logrushook

import (
	"errors"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

type TransErrorLevelLogrusHook struct {
	toLevel          logrus.Level
	excludeCodes     []gocode.Code
	deleteErrorField bool
}

func NewTransErrorLevelLogrusHook(toLevel logrus.Level, excludeCodes []gocode.Code, deleteErrorField bool) logrus.Hook {
	return &TransErrorLevelLogrusHook{
		toLevel:          toLevel,
		excludeCodes:     excludeCodes,
		deleteErrorField: deleteErrorField,
	}
}

func (hook *TransErrorLevelLogrusHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *TransErrorLevelLogrusHook) Fire(entry *logrus.Entry) error {
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

	entry.Level = hook.toLevel

	return nil
}
