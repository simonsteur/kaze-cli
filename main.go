package main

import (
	"flag"
	"os"
)

var (
	//load config
	config = Cfg()

	// api endpoints
	apibase        = "http://" + config.Sensu + ":" + config.Port
	clientapi      = apibase + "/clients"
	checksapi      = apibase + "/checks"
	resultssapi    = apibase + "/results"
	aggregatessapi = apibase + "/aggregates"
	eventssapi     = apibase + "/events"
	silencedsapi   = apibase + "/silenced"
	stashesapi     = apibase + "/stashes"
	healthapi      = apibase + "/health"
	infoapi        = apibase + "/info"

	// generic
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
	silenceClear        bool
	silenceList         bool
	silenceSubscription bool
	// check & resolve command specific
	checkName string
	checkAll  bool
)

func main() {

	// list subcommand
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	// list flags
	listCmd.BoolVar(&client, "client", false, "use to list client(s)")
	listCmd.BoolVar(&check, "check", false, "use to list check(s)")
	listCmd.BoolVar(&event, "event", false, "use to list event(s)")
	listCmd.BoolVar(&silence, "silence", false, "use to list silence entr(y)(ies)")
	listCmd.BoolVar(&result, "result", false, "use to list result(s)")
	listCmd.BoolVar(&aggregate, "aggregate", false, "luse to ist aggregate(s)")
	listCmd.BoolVar(&stash, "stash", false, "use to list stash(es)")
	listCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to list")

	// create subcommand
	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	// create flags
	createCmd.BoolVar(&client, "client", false, "use to create client(s)")
	createCmd.BoolVar(&result, "result", false, "use to create result(s)")
	createCmd.BoolVar(&stash, "stash", false, "use to create stash(es)")
	createCmd.BoolVar(&bulkFile, "file", false, "a valid json file for creation of objects")

	// delete subcommand
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	// delete flags
	deleteCmd.BoolVar(&client, "client", false, "use to delete client(s)")
	deleteCmd.BoolVar(&event, "events", false, "use to delete event(s)")
	deleteCmd.BoolVar(&result, "result", false, "use to delete result(s)")
	deleteCmd.BoolVar(&aggregate, "aggregate", false, "use to delete aggregate(s)")
	deleteCmd.BoolVar(&stash, "stash", false, "use to delete stash(es)")
	deleteCmd.StringVar(&name, "name", "", "specify the name(s) of the object(s) to delete")

	// silence subcommand
	silenceCmd := flag.NewFlagSet("silence", flag.ExitOnError)
	// silence flags
	silenceCmd.BoolVar(&silenceClear, "clear", false, "use to clear silenced entr(y)(ies)")
	silenceCmd.BoolVar(&silenceList, "list", false, "use to list silenced entr(y(ies)")
	silenceCmd.BoolVar(&client, "client", false, "use to target client(s)")
	silenceCmd.BoolVar(&silenceSubscription, "subscription", false, "use to target subscription(s)")
	silenceCmd.StringVar(&name, "name", "", "specify the name(s) of the client(s) or subscription(s)")

	//check subcommand
	checkCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//check flags
	checkCmd.StringVar(&name, "client-name", "", "specify the name of the client")
	checkCmd.StringVar(&checkName, "check-name", "", "specify the name of the check")
	checkCmd.BoolVar(&checkAll, "all", false, "use to target all checks")
	checkCmd.BoolVar(&checkResult, "result", false, "use to get the result back from the requested check")

	//resolve subcommand
	resolveCmd := flag.NewFlagSet("check", flag.ExitOnError)
	//resolve flags
	resolveCmd.StringVar(&name, "client-name", "", "specify the name of the client")
	resolveCmd.StringVar(&checkName, "check-name", "", "specify the name of the check")
	resolveCmd.BoolVar(&checkAll, "all", false, "use to target all events")

	switch os.Args[1] {
	case "list":
		listCmd.Parse(os.Args[2:])
	case "create":
		listCmd.Parse(os.Args[2:])
	case "delete":
		listCmd.Parse(os.Args[2:])
	case "silence":
		listCmd.Parse(os.Args[2:])
	case "check":
		listCmd.Parse(os.Args[2:])
	case "resolve":
		listCmd.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if listCmd.Parsed() {
		cmdController("list")
	}
	if createCmd.Parsed() {
		cmdController("create")
	}
	if deleteCmd.Parsed() {
		cmdController("delete")
	}
	if silenceCmd.Parsed() {
		cmdController("silence")
	}
	if checkCmd.Parsed() {
		cmdController("check")
	}
	if checkCmd.Parsed() {
		cmdController("resolve")
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
