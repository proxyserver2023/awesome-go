package practice

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.WarnLevel)
}

// Run - drives the practice package
func Run() {
	log.WithFields(
		log.Fields{
			"animal": "walrus",
			"size":   10,
		},
	).Info("A group of walrus emerges from the Ocean")

	log.WithFields(
		log.Fields{
			"animal": "walrus",
			"size":   10,
		},
	).Warn("A group of walrus emerges from the Ocean")

	log.WithFields(
		log.Fields{
			"animal": "walrus",
			"size":   10,
		},
	).Fatal("A group of walrus emerges from the Ocean")

	contextLogger := log.WithFields(
		log.Fields{
			"common": "this is a common field",
			"other":  "I also should be logged always",
		},
	)

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")
}
