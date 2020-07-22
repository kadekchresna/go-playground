package other

import log "github.com/sirupsen/logrus"

func Other() {
	log.WithFields(log.Fields{
		"Animal": "Logrus",
	}).Info("A logrus appears")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")
}
