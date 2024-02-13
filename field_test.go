package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithField("username", "Robi").Info("Hello World")
	logger.WithField("username", "Robi Laruku").WithField("name", "Laraku").Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "robi",
		"name": "Robi Laruku",
	}).Info("Hello World")
}
