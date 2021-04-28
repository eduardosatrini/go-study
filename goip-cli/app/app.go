package app

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
)

func searchIP(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println("IP", ip)
	}
}

func searchServer(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println("Server", server.Host)
	}
}

// Generate a app
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line app."
	app.Usage = "Search IP's and namehost."

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search ips on internet.",
			Flags:  flags,
			Action: searchIP,
		},
		{
			Name:   "server",
			Usage:  "Search servers on internet.",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}


