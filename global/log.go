package global

import (
	"io"

	nested "github.com/antonfisher/nested-logrus-formatter"
	"github.com/sirupsen/logrus"
)

func NewLog(w io.Writer, level logrus.Level) *logrus.Logger {
	log := logrus.New()
	log.Out = w
	log.Level = level
	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})
	return log
}
