package main

import (
	log "github.com/sirupsen/logrus"
	"server/router"
)

func init() {
	log.SetReportCaller(true)
}
func main() {
	//router := gin.Default()
	router.JwtEncryption()
	// start

}
