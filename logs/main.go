package logs

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger = logrus.WithFields(logrus.Fields{
	"service": "de-stats",
	"art-id": "de-stats",
	"group": "org.cyverse",
})

func Init(debug *string){
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	if *debug == "on" {
		logrus.SetLevel(logrus.DebugLevel)
		Logger.Debug("Debug mode enabled.")
	}
}