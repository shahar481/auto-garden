package argparsing

import (
	"context"
	"flag"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

type Args struct {
	DbIp *string
	DbPort *int
	DbName *string
	DbUsername *string
	DbPassword *string
	WebListenIp *string
	WebListenPort *int
	WebKeyPath *string
	WebCertPath *string
	WebRootDir *string
	RunTls *bool
	DbConn *pgxpool.Pool
	NoCli *bool
}

var (
	argsInstance *Args
)

// Parses args given to program and puts them into an Args struct
func ParseArgs() *Args {
	if argsInstance != nil {
		return argsInstance
	}
	argsInstance = &Args{}
	argsInstance.DbIp = flag.String("dbip", "127.0.0.1", "Remote database ip")
	argsInstance.DbPort = flag.Int("dbport", 5432, "Remote database port")
	argsInstance.DbName = flag.String("dbname", "postgres", "Remote database name")
	argsInstance.DbUsername = flag.String("dbuser", "postgres", "Remote db username")
	argsInstance.DbPassword = flag.String("dbpass", "Password1", "Remote db crypto")
	argsInstance.WebListenIp = flag.String("weblistenip", "0.0.0.0", "HTTPS server listen ip")
	argsInstance.WebListenPort = flag.Int("weblistenport", 443, "HTTPS server listen port")
	argsInstance.WebKeyPath = flag.String("keypath", "./key", "HTTPS key file")
	argsInstance.WebCertPath = flag.String("certpath", "./cert", "HTTPS cert file")
	argsInstance.RunTls = flag.Bool("ssl", true, "Should run web server in ssl")
	argsInstance.WebRootDir = flag.String("rootdir", "./web/", "Root dir of web")
	argsInstance.NoCli = flag.Bool("nocli", false, "Should run cli")
	flag.Parse()
	argsInstance.connectToDb()
	return argsInstance
}

// Sets the args instance static var.
// WARNING: NOT THREAD SAFE, ONLY RUN BEFORE RUNNING WEB SERVER
func SetArgs(args *Args) {
	argsInstance = args
}

const dbConnectionString = "postgres://%s:%s@%s:%d/%s"


func (a *Args) connectToDb() {
	conn, err := pgxpool.Connect(context.Background(), fmt.Sprintf(dbConnectionString,
		*(a.DbUsername),
		*(a.DbPassword),
		*(a.DbIp),
		*(a.DbPort),
		*(a.DbName)))
	if err != nil {
		log.Fatal(err)
	}
	a.DbConn = conn
}
