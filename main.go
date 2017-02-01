package main

import "github.com/urfave/cli"
import "os"

var (
	config    = Cfg()
	clientapi = "http://" + config.Sensu + ":" + config.Port + "/clients"

	// client flag vars
	clientList          bool
	clientCreate        bool
	clientDelete        bool
	clientName          string
	clientAddress       string
	clientEnvironment   string
	clientSubscriptions []string
)

func main() {

	app := cli.NewApp()
	app.Name = "sensuamplo"
	app.Version = "0.1"
	app.Usage = "control sensu from a cli"
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		{
			Name:  "client",
			Usage: "use to add a client to sensu (most likely a proxy client)",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:        "l, list",
					Usage:       "list clients",
					Destination: &clientList,
				},
				cli.BoolFlag{
					Name:        "c, create",
					Usage:       "create clients",
					Destination: &clientCreate,
				},
				cli.BoolFlag{
					Name:        "d, delete",
					Usage:       "delete clients",
					Destination: &clientDelete,
				},
				cli.StringFlag{
					Name:        "name",
					Usage:       "name of the client",
					Destination: &clientName,
				},
				cli.StringFlag{
					Name:        "environment, env",
					Usage:       "address of the client",
					Destination: &clientEnvironment,
				},
				cli.StringFlag{
					Name:        "address",
					Usage:       "address of the client",
					Destination: &clientAddress,
				},
				cli.StringSliceFlag{
					Name:  "subscriptions",
					Usage: "address of the client",
				},
			},
		},
	}
	app.Run(os.Args)
}
