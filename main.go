package main

import (
	"flag"
	"os"
)

var (
	config    = Cfg()
	apibase   = "http://" + config.Sensu + ":" + config.Port
	clientapi = apibase + "/clients"
	checksapi = apibase + "/checks"

	// client flag vars
	client              bool
	checks              bool
	events              bool
	silence             bool
	results             bool
	name                string
	clientBulk          bool
	clientBulkFile      string
	clientAddress       string
	clientEnvironment   string
	clientSubscriptions []string
)

func main() {

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.BoolVar(&client, "client", false, "specify to list clients")
	listCmd.BoolVar(&checks, "checks", false, "specify to list checks")
	listCmd.BoolVar(&events, "events", false, "specify to list events")

	// error message

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCmd.Parsed() {
		handleListCmd(client, checks, events)
	}

	// app := cli.NewApp()
	// app.Name = "kaze"
	// app.Version = "1.0"
	// app.Usage = "control sensu from a cli"
	// app.EnableBashCompletion = true

	// app.Commands = []cli.Command{
	// 	{
	// 		Name:  "list",
	// 		Usage: "use to add a client to sensu (most likely a proxy client)",
	// 		Flags: []cli.Flag{
	// 			cli.BoolFlag{
	// 				Name:        "l, list",
	// 				Usage:       "list clients",
	// 				Destination: &clientList,
	// 			},
	// 			cli.BoolFlag{
	// 				Name:        "c, create",
	// 				Usage:       "create clients",
	// 				Destination: &clientCreate,
	// 			},
	// 			cli.BoolFlag{
	// 				Name:        "d, delete",
	// 				Usage:       "delete clients",
	// 				Destination: &clientDelete,
	// 			},
	// 			cli.StringFlag{
	// 				Name:        "f, file",
	// 				Usage:       "specify when creating clients to do a bulk operation. Has to be a correctly formatted json file.",
	// 				Destination: &clientBulkFile,
	// 			},
	// 			cli.StringFlag{
	// 				Name:        "name",
	// 				Usage:       "name of the client (required for create)",
	// 				Destination: &clientName,
	// 			},
	// 			cli.StringFlag{
	// 				Name:        "environment, env",
	// 				Usage:       "environment of the client (required for create)",
	// 				Destination: &clientEnvironment,
	// 			},
	// 			cli.StringFlag{
	// 				Name:        "address",
	// 				Usage:       "address of the client (required for create)",
	// 				Destination: &clientAddress,
	// 			},
	// 			cli.StringSliceFlag{
	// 				Name:  "subscriptions",
	// 				Usage: "subcriptions of the client (required for create)",
	// 			},
	// 		},
	// 		Action: manageClient,
	// 	},
	// }
	// app.Run(os.Args)
}
