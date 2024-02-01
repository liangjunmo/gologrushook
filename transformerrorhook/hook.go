package transformerrorhook

import (
	"errors"

	"github.com/liangjunmo/gocode"
	"github.com/sirupsen/logrus"
)

type Hook struct {
	toLevel        logrus.Level
	excludedCodes  []gocode.Code
	deleteErrorKey bool
}

func New(toLevel logrus.Level) *Hook {
	return &Hook{
		toLevel: toLevel,
	}
}

func (hook *Hook) ExcludedCodes(codes []gocode.Code) {
	hook.excludedCodes = codes
}

func (hook *Hook) DeleteErrorKey() {
	hook.deleteErrorKey = true
}

func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (hook *Hook) Fire(entry *logrus.Entry) error {
	err, ok := entry.Data[logrus.ErrorKey].(error)
	if !ok || err == nil {
		return nil
	}

	if hook.deleteErrorKey {
		delete(entry.Data, logrus.ErrorKey)
	}

	for _, code := range hook.excludedCodes {
		if errors.Is(err, code) {
			return nil
		}
	}

	entry.Level = hook.toLevel

	return nil
}
