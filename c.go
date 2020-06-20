package module_c

import (
	log "github.com/sirupsen/logrus"
)

var VersionConfirm = "v1.0.0"

func init() {
	log.Infof("Version %v load", VersionConfirm)
}

func GetVersionConfirm() string {
	return VersionConfirm
}

func SetVersionConfirm(s string) {
	VersionConfirm = s
}
