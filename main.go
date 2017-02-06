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
	clientBulk          bool
	clientBulkFile      string
	clientName          string
	clientAddress       string
	clientEnvironment   string
	clientSubscriptions []string
)

func main() {

	app := cli.NewApp()
	app.Name = "kaze"
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
				cli.BoolFlag{
					Name:        "b, bulk",
					Usage:       "use if you wish to create or delete clients in bulk, requires the --file flag",
					Destination: &clientBulk,
				},
				cli.StringFlag{
					Name:        "f, file",
					Usage:       "required when using bylk & delete. Has to be a correctly formatted json file.",
					Destination: &clientBulkFile,
				},
				cli.StringFlag{
					Name:        "name",
					Usage:       "name of the client (required for create)",
					Destination: &clientName,
				},
				cli.StringFlag{
					Name:        "environment, env",
					Usage:       "environment of the client (required for create)",
					Destination: &clientEnvironment,
				},
				cli.StringFlag{
					Name:        "address",
					Usage:       "address of the client (required for create)",
					Destination: &clientAddress,
				},
				cli.StringSliceFlag{
					Name:  "subscriptions",
					Usage: "subcriptions of the client (required for create)",
				},
			},
			Action: manageClient,
		},
	}
	app.Run(os.Args)
}
