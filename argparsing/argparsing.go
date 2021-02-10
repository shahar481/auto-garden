package argparsing

import "flag"

type Args struct {
	DbIp *string
	DbPort *int
	DbUsername *string
	DbPassword *string
	WebListenIp *string
	WebListenPort *int
	WebKeyPath *string
	WebCertPath *string
	WebRootDir *string
}

// Parses args given to program and puts them into an Args struct
func ParseArgs() Args {
	a := Args{}
	a.DbIp = flag.String("dbip", "127.0.0.1", "Remote database ip")
	a.DbPort = flag.Int("dbport", 5432, "Remote database port")
	a.DbUsername = flag.String("dbuser", "postgres", "Remote db username")
	a.DbPassword = flag.String("dbpass", "Password1", "Remote db password")
	a.WebListenIp = flag.String("weblistenip", "0.0.0.0", "HTTPS server listen ip")
	a.WebListenPort = flag.Int("weblistenport", 443, "HTTPS server listen port")
	a.WebKeyPath = flag.String("keypath", "./key", "HTTPS key file")
	a.WebCertPath = flag.String("certpath", "./cert", "HTTPS cert file")
	a.WebRootDir = flag.String("rootdir", "./web/", "Root dir of web")
	flag.Parse()
	return a
}
