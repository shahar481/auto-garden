package testutility

import (
	"auto-garden/argparsing"
)

func CreateArgs(dbIp string, dbPort int, dbUsername string, dbPassword string, dbName string,
	webListenIp string, webListenPort int, webKeyPath string, webCertPath string, runTls bool,
	webRootDir string) *argparsing.Args {
	args := &argparsing.Args{}
	args.DbIp = &(dbIp)
	args.DbPort = &(dbPort)
	args.DbUsername = &(dbUsername)
	args.DbPassword = &(dbPassword)
	args.DbName = &(dbName)
	args.WebListenIp = &(webListenIp)
	args.WebListenPort = &(webListenPort)
	args.WebKeyPath = &(webKeyPath)
	args.WebCertPath = &(webCertPath)
	args.RunTls = &(runTls)
	args.WebRootDir = &(webRootDir)
	return args
}