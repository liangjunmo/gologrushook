package logrushook

import (
	"errors"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

type TransformErrorLevelLogrusHook struct {
	toLevel        logrus.Level
	excludeCodes   []gocode.Code
	deleteErrorKey bool
}

func NewTransformErrorLevelLogrusHook(toLevel logrus.Level, excludeCodes []gocode.Code, deleteErrorField bool) logrus.Hook {
	return &TransformErrorLevelLogrusHook{
		toLevel:        toLevel,
		excludeCodes:   excludeCodes,
		deleteErrorKey: deleteErrorField,
	}
}

func (hook *TransformErrorLevelLogrusHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *TransformErrorLevelLogrusHook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok || err == nil {
		return nil
	}

	if hook.deleteErrorKey {
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
