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
	client    bool
	checks    bool
	events    bool
	silence   bool
	results   bool
	aggregate bool
	stash     bool
	name      []string
	bulkFile  string
	// silence command specific
	listSilenced         bool
	clearSilenced        bool
	subscriptionSilenced bool
)

func main() {

	//list command
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.BoolVar(&client, "client", false, "use to list client(s)")
	listCmd.BoolVar(&check, "checks", false, "use to list check(s)")
	listCmd.BoolVar(&event, "events", false, "use to list event(s)")
	listCmd.BoolVar(&silence, "silence", false, "use to list silence entr(y)(ies)")
	listCmd.BoolVar(&result, "result", false, "use to list result(s)")
	listCmd.BoolVar(&aggregate, "aggregate", false, "luse to ist aggregate(s)")
	istCmd.BoolVar(&stash, "stash", false, "use to list stash(es)")
	listCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to list")

	//create command
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	createCmd.BoolVar(&client, "client", false, "use to create client(s)")
	createCmd.BoolVar(&result, "result", false, "use to create result(s)")
	createCmd.BoolVar(&stash, "stash", false, "use to create stash(es)")
	createCmd.BoolVar(&bulkFile, "file", false, "a valid json file for creation of objects")

	//delete command
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	listCmd.BoolVar(&client, "client", false, "use to delete client(s)")
	listCmd.BoolVar(&event, "events", false, "use to delete event(s)")
	listCmd.BoolVar(&result, "result", false, "use to delete result(s)")
	listCmd.BoolVar(&aggregate, "aggregate", false, "use to delete aggregate(s)")
	istCmd.BoolVar(&stash, "stash", false, "use to delete stash(es)")
	listCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to delete")

	//silence command
	silenceCmd := flag.NewFlagSet("silence", flag.ExitOnError)
	listCmd.BoolVar(&clear, "clear", false, "use to clear silenced entr(y)(ies)")
	listCmd.BoolVar(&list, "list", false, "use to list silenced entr(y(ies)")
	listCmd.BoolVar(&client, "client", false, "use to target client(s)")
	listCmd.BoolVar(&check, "subscription", false, "use to target subscription(s)")
	listCmd.StringVar(&name, "name", "", "specify the name(s) of the client(s) or subscription(s)")

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
	case "create":
		listCmd.Parse(os.Args[2:])
	case "delete":
		listCmd.Parse(os.Args[2:])
	case "silence":
		listCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCmd.Parsed() {
		handleListCmd(client, check, event, silence, result, aggregate, stash)
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
