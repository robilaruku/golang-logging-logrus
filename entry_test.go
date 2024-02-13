package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func Test(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "robi")
	entry.Info("Hello Entry")
}
