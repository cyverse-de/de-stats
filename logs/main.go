package logs

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.WithFields(logrus.Fields{
	"service": "de-stats",
	"art-id": "de-stats",
	"group": "org.cyverse",
})

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

}