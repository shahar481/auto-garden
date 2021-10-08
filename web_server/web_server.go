package web_server

import (
	"auto-garden/argparsing"
	"auto-garden/web_server/api"
	"auto-garden/web_server/api/plants"
	"crypto/tls"
	"fmt"
	"github.com/savsgio/atreugo/v11"
	"log"
	"net"
)

var (
	ApiFunctions = map[string]apiFunction {
		"/garden/should-water": {
			Func:       plants.ShouldWaterRequest,
			Parameters: plants.ShouldWaterParameters,
		},
	}
)

// Creates a net.Listener object for tls connections
func createTlsListener(keyPath string, certPath string, listenIp string, listenPort int) (net.Listener, error) {
	cer, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}
	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}
	ln, err := tls.Listen("tcp", fmt.Sprintf("%s:%d",listenIp, listenPort), tlsConfig)
	if err != nil {
		return nil, err
	}
	return ln, nil
}

func createNonTlsListener(listenIp string, listenPort int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%d",listenIp, listenPort))
}

func createListener(args *argparsing.Args) net.Listener {
	var listener net.Listener
	var err error
	if *(args.RunTls) {
		listener, err = createTlsListener(
			*(args.WebKeyPath),
			*(args.WebCertPath),
			*(args.WebListenIp),
			*(args.WebListenPort))
	} else {
		listener, err = createNonTlsListener(
			*(args.WebListenIp),
			*(args.WebListenPort))
	}
	if err != nil {
		log.Fatal(err)
	}
	return listener
}

func StartServer(args *argparsing.Args) {
	config := atreugo.Config{}
	server := atreugo.New(config)

	listener := createListener(args)

	defer listener.Close()

	server.Static("/static", *(args.WebRootDir))
	api.SetHTTPFunctions(server)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}

func setApiFunctions(server *atreugo.Atreugo) {

}