package main

import (
	"fmt"
	"flag"
	"github.com/benmj87/cheap-ftp-server/src/ftpserver"
	"github.com/goftp/server"
)

func main() {
	port := flag.Int("port", 21, "Port")
	hostname := flag.String("hostip", "127.0.0.1", "HostIP to listen on")
	root := flag.String("root", "", "Root folder to write to")
	flag.Parse()

	fact := ftpserver.NewDriverFactory(*root)	
	opts := &server.ServerOpts{
		Factory: fact,
		Port: *port,
		Hostname: *hostname,
		Auth: ftpserver.NewAuth(),
		PublicIp: *hostname,
	}
	
	fmt.Printf("Starting to listen on %v:%v....\n", opts.Hostname, opts.Port)
	server := server.NewServer(opts)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}